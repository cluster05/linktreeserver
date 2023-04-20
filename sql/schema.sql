create table if not exists linktree.auths
(
    authId    varchar(25)  not null
    primary key,
    username  varchar(10)  not null,
    firstname varchar(25)  not null,
    lastname  varchar(25)  not null,
    email     varchar(50)  not null,
    password  varchar(100) not null,
    authMode  varchar(10)  not null,
    createdAt int          not null,
    updatedAt int          not null,
    isDeleted tinyint(1)   not null
    );

create table if not exists linktree.links
(
    authId    varchar(25)  not null,
    linkId    varchar(25)  not null
    primary key,
    title     varchar(40)  not null,
    url       varchar(255) not null,
    imageUrl  varchar(255) not null,
    createdAt int          not null,
    updatedAt int          not null,
    isDeleted int          not null,
    constraint links_auth_fk
    foreign key (authId) references linktree.auths (authId)
    );

create table if not exists linktree.analytics
(
    analyticsId   varchar(25) not null
    primary key,
    linkId        varchar(25) not null,
    continentCode varchar(5)  not null,
    countryCode   varchar(5)  not null,
    regionCode    varchar(5)  not null,
    city          varchar(35) not null,
    pincode       varchar(8)  not null,
    latitude      varchar(10) not null,
    longitude     varchar(10) not null,
    userAgent     varchar(10) not null,
    os            varchar(10) not null,
    createdAt     int         not null,
    updatedAt     int         not null,
    constraint analytics_link_fk
    foreign key (linkId) references linktree.links (linkId)
    );

create table if not exists linktree.payments
(
    paymentId varchar(25) not null
    primary key,
    authId    varchar(25) not null,
    constraint payments_auth_fk
    foreign key (authId) references linktree.auths (authId)
    );

create table if not exists linktree.plans
(
    planId           varchar(25) not null
    primary key,
    authId           varchar(25) not null,
    paymentId        varchar(25) not null,
    planType         varchar(10) not null,
    start            int         not null,
    end              int         not null,
    subscriptionType varchar(10) not null,
    createAt         int         not null,
    updatedAt        int         not null,
    activeStatus     tinyint(1)  not null,
    constraint plans_auth_fk
    foreign key (authId) references linktree.auths (authId),
    constraint plans_payment_fk
    foreign key (paymentId) references linktree.payments (paymentId)
    );

