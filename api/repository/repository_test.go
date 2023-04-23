package repository_test

import (
	"fmt"
	"github.com/cluster05/linktree/api/model"
	"github.com/cluster05/linktree/api/repository"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var (
	linktreeDB = "linktree.db"
)

var _ = Describe("Repository", func() {

	Context("Auth Repository", Ordered, func() {

		var (
			testUsername    = "test-username"
			testEmail       = "email@test.com"
			testPassword    = "test-password"
			testNewPassword = "test-new-password"
		)

		var (
			repo repository.AuthRepository
			db   *gorm.DB
			err  error
		)

		BeforeEach(func() {
			db, err = gorm.Open(sqlite.Open(linktreeDB), &gorm.Config{})
			if err != nil {
				panic("failed to connect database")
			}
			repo = repository.NewAuthRepository(&repository.AuthRepositoryConfig{
				MySqlDB: db,
			})
			err = db.AutoMigrate(&model.Auth{})
			if err != nil {
				panic("failed to migrate tables")
			}
		})

		AfterEach(func() {
			_ = db.Migrator().DropTable(&model.Auth{})
		})

		AfterAll(func() {
			err := os.Remove(linktreeDB)
			if err != nil {
				fmt.Println("error while clearing database", linktreeDB, err)
			}
		})

		Describe("FetchAuthByUsername", func() {
			It("should return a user with the given username", func() {
				expectedAuth := model.Auth{
					Username: testUsername,
					Password: testPassword,
					Email:    testEmail,
				}
				result := db.Create(&expectedAuth)
				if result.Error != nil {
					panic("failed to create test data")
				}

				auth, err := repo.FetchAuthByUsername(expectedAuth.Username)
				Expect(err).To(BeNil())
				Expect(auth.Username).To(Equal(expectedAuth.Username))
				Expect(auth.Password).To(Equal(expectedAuth.Password))
				Expect(auth.Email).To(Equal(expectedAuth.Email))
			})
		})

		Describe("FetchAuthByEmail", func() {
			It("should return a user with the given email", func() {
				expectedAuth := model.Auth{
					Username: testUsername,
					Password: testPassword,
					Email:    testEmail,
				}
				result := db.Create(&expectedAuth)
				if result.Error != nil {
					panic("failed to create test data")
				}

				auth, err := repo.FetchAuthByEmail(expectedAuth.Email)
				Expect(err).To(BeNil())
				Expect(auth.Username).To(Equal(expectedAuth.Username))
				Expect(auth.Password).To(Equal(expectedAuth.Password))
				Expect(auth.Email).To(Equal(expectedAuth.Email))
			})
		})

		Describe("Register", func() {
			It("should create and return a new user", func() {
				record := model.Auth{
					Username: testUsername,
					Password: testPassword,
					Email:    testEmail,
				}
				auth, err := repo.Register(record)
				Expect(err).To(BeNil())
				Expect(auth.Username).To(Equal(record.Username))
				Expect(auth.Password).To(Equal(record.Password))
				Expect(auth.Email).To(Equal(record.Email))

				var count int64
				db.Model(&model.Auth{}).Count(&count)
				Expect(count).To(Equal(int64(1)))
			})
		})

		Describe("ChangePassword", func() {
			It("should update the password of the user with the given authId", func() {
				record := model.Auth{
					Username: testUsername,
					Password: testPassword,
					Email:    testEmail,
				}
				result := db.Create(&record)
				if result.Error != nil {
					panic("failed to create test data")
				}

				err := repo.ChangePassword(testNewPassword, record.AuthId)
				Expect(err).To(BeNil())

				updatedAuth := model.Auth{}
				db.Where("authId = ?", record.AuthId).First(&updatedAuth)
				Expect(updatedAuth.Password).To(Equal(testNewPassword))
			})
		})

	})

	Context("Link Repository", Ordered, func() {

		var (
			testLinkId   = "test_link_id"
			testAuthId   = "test_auth_id"
			testTitle    = "Test Link"
			testURL      = "https://example.com"
			testImageURL = "https://example.com/image.jpg"

			testInvalidLinkId = "invalid_link_id"
		)

		var (
			repo repository.LinkRepository
			db   *gorm.DB
			err  error
		)

		BeforeEach(func() {
			db, err = gorm.Open(sqlite.Open(linktreeDB), &gorm.Config{})
			if err != nil {
				panic("failed to connect database")
			}
			repo = repository.NewLinkRepository(&repository.LinkRepositoryConfig{
				MySqlDB: db,
			})
			err = db.AutoMigrate(&model.Link{})
			if err != nil {
				panic("failed to migrate tables")
			}
		})

		AfterEach(func() {
			_ = db.Migrator().DropTable(&model.Link{})
		})

		AfterAll(func() {
			err := os.Remove(linktreeDB)
			if err != nil {
				fmt.Println("error while clearing database", linktreeDB, err)
			}
		})

		Context("CreateLink", func() {
			It("should create a new link", func() {
				link := model.Link{
					LinkId: testLinkId,
					AuthId: testAuthId,
					Title:  testTitle,
					URL:    testURL,
				}
				createdLink, err := repo.CreateLink(link)
				Expect(err).Should(BeNil())

				Expect(createdLink.LinkId).ShouldNot(BeEmpty())
			})
		})

		Context("ReadLink", func() {
			var (
				link1 model.Link
				link2 model.Link
			)
			BeforeEach(func() {
				link1 = model.Link{
					AuthId:   testAuthId,
					LinkId:   testLinkId,
					Title:    testTitle + " 1",
					URL:      testURL + " 1",
					ImageURL: testImageURL + " 1",
				}
				link2 = model.Link{
					AuthId:   testAuthId,
					LinkId:   testLinkId,
					Title:    testTitle + " 2",
					URL:      testURL + " 2",
					ImageURL: testImageURL + " 2",
				}
				_, err := repo.CreateLink(link1)
				Expect(err).NotTo(HaveOccurred())
				_, err = repo.CreateLink(link2)
				Expect(err).NotTo(HaveOccurred())
			})

			It("should read all links for an auth ID", func() {
				links, err := repo.ReadLink(testAuthId)

				Expect(err).NotTo(HaveOccurred())
				Expect(links).To(HaveLen(2))

			})
		})

		Context("FindLink", func() {
			var (
				link model.Link
			)
			BeforeEach(func() {
				link = model.Link{
					LinkId:   testLinkId,
					AuthId:   testAuthId,
					Title:    testTitle,
					URL:      testURL,
					ImageURL: testImageURL,
				}
				createdLink, err := repo.CreateLink(link)
				Expect(err).NotTo(HaveOccurred())
				link = createdLink
			})

			It("should find a link by auth ID and link ID", func() {
				foundLink, err := repo.FindLink(link.AuthId, link.LinkId)

				Expect(err).NotTo(HaveOccurred())
				Expect(foundLink).To(Equal(link))
			})

			It("should return an error if the link is not found", func() {
				_, err := repo.FindLink("invalid_auth_id", "invalid_link_id")

				Expect(err).To(MatchError(gorm.ErrRecordNotFound))
			})
		})

		Context("UpdateLink", func() {
			var (
				link model.Link
			)
			BeforeEach(func() {
				link = model.Link{
					LinkId:   testLinkId,
					AuthId:   testAuthId,
					Title:    testTitle,
					URL:      testURL,
					ImageURL: testImageURL,
				}
				createdLink, err := repo.CreateLink(link)
				Expect(err).NotTo(HaveOccurred())
				link = createdLink
			})

			It("should update a link", func() {
				link.Title = testTitle + " updated"
				link.URL = testURL + " updated"
				link.ImageURL = testImageURL + " updated"
				updatedLink, err := repo.UpdateLink(link)

				Expect(err).NotTo(HaveOccurred())
				Expect(updatedLink).To(Equal(link))

				foundLink, err := repo.FindLink(link.AuthId, link.LinkId)
				Expect(err).NotTo(HaveOccurred())
				Expect(foundLink).To(Equal(link))
			})

			It("should return an error if the link is not found", func() {
				link.LinkId = testInvalidLinkId
				_, err := repo.UpdateLink(link)
				Expect(err).To(MatchError(gorm.ErrRecordNotFound))
			})
		})

		Context("DeleteLink", func() {
			var (
				link model.Link
			)
			BeforeEach(func() {
				link = model.Link{
					LinkId:   testLinkId,
					AuthId:   testAuthId,
					Title:    testTitle,
					URL:      testURL,
					ImageURL: testImageURL,
				}
				createdLink, err := repo.CreateLink(link)
				Expect(err).NotTo(HaveOccurred())
				link = createdLink
			})

			It("should delete a link", func() {
				err := repo.DeleteLink(link.LinkId)
				Expect(err).NotTo(HaveOccurred())

				_, err = repo.FindLink(link.AuthId, link.LinkId)
				Expect(err).To(MatchError(gorm.ErrRecordNotFound))
			})

			It("should return an error if the link is not found", func() {
				err := repo.DeleteLink(testInvalidLinkId)

				Expect(err).To(MatchError(gorm.ErrRecordNotFound))
			})
		})

	})

	Context("analytics Repository", Ordered, func() {

		var (
			authRepo      repository.AuthRepository
			linkRepo      repository.LinkRepository
			analyticsRepo repository.AnalyticsRepository
			db            *gorm.DB
			err           error
		)

		var (
			auth model.Auth
			link model.Link
		)

		BeforeEach(func() {
			db, err = gorm.Open(sqlite.Open(linktreeDB), &gorm.Config{})
			if err != nil {
				panic("failed to connect database")
			}

			authRepo = repository.NewAuthRepository(&repository.AuthRepositoryConfig{
				MySqlDB: db,
			})

			linkRepo = repository.NewLinkRepository(&repository.LinkRepositoryConfig{
				MySqlDB: db,
			})

			analyticsRepo = repository.NewAnalyticsRepository(&repository.AnalyticsRepositoryConfig{
				MySqlDB: db,
			})

			err = db.AutoMigrate(&model.Auth{})
			err = db.AutoMigrate(&model.Link{})
			err = db.AutoMigrate(&model.Analytics{})
			if err != nil {
				panic("failed to migrate tables")
			}

			auth = model.Auth{}
			auth, err := authRepo.Register(auth)
			Expect(err).NotTo(HaveOccurred())

			link = model.Link{
				AuthId: auth.AuthId,
			}
			link, err = linkRepo.CreateLink(link)
			Expect(err).NotTo(HaveOccurred())

		})

		AfterEach(func() {
			_ = db.Migrator().DropTable(&model.Analytics{})
			_ = db.Migrator().DropTable(&model.Link{})
			_ = db.Migrator().DropTable(&model.Auth{})
		})

		AfterAll(func() {
			err := os.Remove(linktreeDB)
			if err != nil {
				fmt.Println("error while clearing database", linktreeDB, err)
			}
		})

		Describe("CreateAnalytics", func() {
			It("should create new analytics and return it ", func() {

				analytics := model.Analytics{
					LinkId: link.LinkId,
				}
				createdAnalytic, err := analyticsRepo.CreateAnalytics(analytics)
				Expect(err).Should(BeNil())

				Expect(createdAnalytic).Should(BeTrue())
			})
		})

		Describe("Read Analytics", func() {

			var (
				analytics1 model.Analytics
				analytics2 model.Analytics
			)

			BeforeEach(func() {

				analytics1 = model.Analytics{
					LinkId: link.LinkId,
				}

				analytics2 = model.Analytics{
					LinkId: link.LinkId,
				}

				_, err = analyticsRepo.CreateAnalytics(analytics1)
				Expect(err).NotTo(HaveOccurred())
				_, err = analyticsRepo.CreateAnalytics(analytics2)
				Expect(err).NotTo(HaveOccurred())
			})

			It("should read all analytics for an link ID", func() {

				user := model.JWTPayload{AuthId: auth.AuthId}

				_, err := analyticsRepo.ReadAnalytics(user)
				Expect(err).NotTo(HaveOccurred())

			})
		})
	})
})
