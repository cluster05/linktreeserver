package service_test

import (
	"time"

	"github.com/cluster05/linktree/api/config"
	"github.com/cluster05/linktree/api/model"
	"github.com/cluster05/linktree/api/service"
	"github.com/dgrijalva/jwt-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {

	Context("JWT service", func() {

		Describe("GenerateToken", func() {
			var (
				auth       model.Auth
				jwtPayload model.JWTPayload
			)

			BeforeEach(func() {
				auth = model.Auth{
					AuthId:    "1",
					Username:  "testuser",
					Firstname: "Test",
					Lastname:  "User",
					Email:     "testuser@example.com",
					AuthMode:  "email",
				}

				config.AppConfig.JWTSecret = "test-secret"
				config.AppConfig.TokenExpireDuration = 15

				now := time.Now().Unix()
				exp := time.Now().Add(time.Duration(config.AppConfig.TokenExpireDuration) * time.Minute).Unix()

				jwtPayload = model.JWTPayload{
					AuthId:    auth.AuthId,
					Username:  auth.Username,
					Firstname: auth.Firstname,
					Lastname:  auth.Lastname,
					Email:     auth.Email,
					AuthMode:  auth.AuthMode,
					IssuedAt:  now,
					ExpiredAt: exp,
				}
			})

			It("should generate a JWT token with correct claims", func() {
				token, err := service.GenerateToken(auth)
				Expect(err).To(BeNil())
				Expect(token).NotTo(BeEmpty())

				parsedToken, err := jwt.ParseWithClaims(token, &model.JWTPayload{}, func(token *jwt.Token) (interface{}, error) {
					return []byte(config.AppConfig.JWTSecret), nil
				})

				Expect(err).To(BeNil())
				Expect(parsedToken.Valid).To(BeTrue())

				payload, ok := parsedToken.Claims.(*model.JWTPayload)
				Expect(ok).To(BeTrue())
				Expect(payload).To(Equal(&jwtPayload))
			})
		})

	})

	Context("Auth service", func() {
		/*
			Describe("AuthService", func() {
				var (
					authRepository *repository.MockAuthRepository
					authService    service.AuthService
					jwtSecret      = "test-secret"
				)

				BeforeEach(func() {
					authRepository = &repository.MockAuthRepository{}
					authService = service.NewAuthService(&service.AuthServiceConfig{
						AuthRepository: authRepository,
					})
				})

				Describe("Register", func() {
					var (
						registerDTO model.RegisterDTO
						hashedPwd   string
						auth        model.Auth
						jwtToken    string
						err         error
					)

					BeforeEach(func() {
						registerDTO = model.RegisterDTO{
							Username:  "username",
							Firstname: "Test",
							Lastname:  "User",
							Email:     "test@example.com",
							Password:  "password",
						}

						hashedPwd, _ = hash.CreatePasswordHash(registerDTO.Password)

						auth = model.Auth{
							Username:  registerDTO.Username,
							Firstname: registerDTO.Firstname,
							Lastname:  registerDTO.Lastname,
							Email:     registerDTO.Email,
							Password:  hashedPwd,
							AuthMode:  model.EmailAuthMode,
						}

						authRepository.On("FetchAuthByEmail", registerDTO.Email).Return(nil, gorm.ErrRecordNotFound)
						authRepository.On("FetchAuthByUsername", registerDTO.Username).Return(nil, gorm.ErrRecordNotFound)
						authRepository.On("Register", auth).Return(auth, nil)
						jwtToken, err = service.GenerateToken(auth)
					})

					It("should successfully register a new account and generate a JWT token", func() {
						token, err := authService.Register(registerDTO)
						Expect(err).To(BeNil())
						Expect(token).NotTo(BeEmpty())
						Expect(jwtToken).To(Equal(token))
					})

					It("should return an error if an account already exists with the same email address", func() {
						authRepository.On("FetchAuthByEmail", registerDTO.Email).Return(&model.Auth{}, nil)

						token, err := authService.Register(registerDTO)
						assert.Empty(GinkgoT(), token)
						assert.Equal(GinkgoT(), service.ErrAccountAlreadyExists, err)
					})

					It("should return an error if an account already exists with the same username", func() {
						authRepository.On("FetchAuthByUsername", registerDTO.Username).Return(&model.Auth{}, nil)

						token, err := authService.Register(registerDTO)
						assert.Empty(GinkgoT(), token)
						assert.Equal(GinkgoT(), service.ErrUsernameAlreadyExists, err)
					})

					It("should return an error if there's an error while registering the user", func() {
						authRepository.On("Register", auth).Return(model.Auth{}, errors.New("database error"))

						token, err := authService.Register(registerDTO)
						assert.Empty(GinkgoT(), token)
						assert.NotNil(GinkgoT(), err)
					})
				})

				Describe("Login", func() {
					var (
						loginDTO     model.LoginDTO
						hashedPwd    string
						auth         model.Auth
						jwtToken     string
						expectedErr  error
						returnedErr  error
						returnedAuth *model.Auth
						returnedJwt  string
					)

					BeforeEach(func() {
						loginDTO = model.LoginDTO{
							Email:    "test@example.com",
							Password: "password",
						}

						hashedPwd, _ = hash.CreatePasswordHash(loginDTO.Password)

						auth = model.Auth{
							Username:  "username",
							Firstname: "Test",
							Lastname:  "User",
							Email:     loginDTO.Email,
							Password:  hashedPwd,
							AuthMode:  model.EmailAuthMode,
						}

						jwtToken, _ = service.GenerateToken(auth)

						expectedErr = nil
						returnedErr = nil
						returnedAuth = &auth
						returnedJwt = jwtToken

						authRepository.On("FetchAuthByEmail", loginDTO.Email).Return(returnedAuth, returnedErr)
					})

					It("should successfully log in an existing user and generate a JWT token", func() {
						token, err := authService.Login(loginDTO)
						Expect(err).To(BeNil())
						Expect(token).NotTo(BeEmpty())
						Expect(jwtToken).To(Equal(token))
					})

					It("should return an error if there's no account associated with the provided email address", func() {
						authRepository.On("FetchAuthByEmail", loginDTO.Email).Return(nil, gorm.ErrRecordNotFound)

						token, err := authService.Login(loginDTO)
						assert.Empty(GinkgoT(), token)
						assert.Equal(GinkgoT(), service.ErrorAccountNotExists, err)
					})

					It("should return an error if the user provides invalid credentials", func() {
						returnedAuth.Password = "invalid-password"
						authRepository.On("FetchAuthByEmail", loginDTO.Email).Return(returnedAuth, nil)

						token, err := authService.Login(loginDTO)
						assert.Empty(GinkgoT(), token)
						assert.Equal(GinkgoT(), service.ErrorInvalidCredentials, err)
					})
				})

				Describe("ChangePassword", func() {
					var (
						changePasswordDTO model.ChangePasswordDTO
						hashedOldPwd      string
						hashedNewPwd      string
						auth              model.Auth
						jwtToken          string
						expectedErr       error
						returnedErr       error
						returnedAuth      *model.Auth
						returnedJwt       string
					)

					BeforeEach(func() {
						changePasswordDTO = model.ChangePasswordDTO{
							Email:           "test@example.com",
							OldPassword:     "password",
							NewPassword:     "new-password",
							ConfirmPassword: "new-password",
						}

						hashedOldPwd, _ = hash.CreatePasswordHash(changePasswordDTO.OldPassword)
						hashedNewPwd, _ = hash.CreatePasswordHash(changePasswordDTO.NewPassword)

						auth = model.Auth{
							Username:  "username",
							Firstname: "Test",
							Lastname:  "User",
							Email:     changePasswordDTO.Email,
							Password:  hashedOldPwd,
							AuthMode:  model.EmailAuthMode,
						}

						jwtToken, _ = service.GenerateToken(auth)

						expectedErr = nil
						returnedErr = nil
						returnedAuth = &auth
						returnedJwt = jwtToken

						authRepository.On("FetchAuthByEmail", changePasswordDTO.Email).Return(returnedAuth, returnedErr)
						authRepository.On("ChangePassword", hashedNewPwd, auth.AuthId).Return(nil)
					})

					It("should successfully change the password for an existing user and generate a new JWT token", func() {
						token, err := authService.ChangePassword(changePasswordDTO)
						Expect(err).To(BeNil())
						Expect(token).NotTo(BeEmpty())
						assert.NotEqual(GinkgoT(), jwtToken, token)
					})

					It("should return an error if there's no account associated with the provided email address", func() {
						authRepository.On("FetchAuthByEmail", changePasswordDTO.Email).Return(nil, gorm.ErrRecordNotFound)

						token, err := authService.ChangePassword(changePasswordDTO)
						assert.Empty(GinkgoT(), token)
						assert.Equal(GinkgoT(), service.ErrorAccountNotExists, err)
					})

					It("should return an error if the user provides invalid old password", func() {
						returnedAuth.Password = "invalid-password"
						authRepository.On("FetchAuthByEmail", changePasswordDTO.Email).Return(returnedAuth, nil)

						token, err := authService.ChangePassword(changePasswordDTO)
						assert.Empty(GinkgoT(), token)
						assert.Equal(GinkgoT(), service.ErrorInvalidCredentials, err)
					})
				})
			})
		*/
	})

	Context("Analytics service", func() {})

	Context("Link service", func() {})

	Context("Plan service", func() {})

	Context("Plan service", func() {})
})
