![goecs](img/goecs.png)

A distributed microservice e-commerce system based on Hertz and gRPC.

## Structure

```sh
├── LICENSE
├── README.md
├── apis
│   ├── go.mod
│   ├── go.sum
│   ├── goods-web
│   │   ├── api
│   │   │   ├── banners
│   │   │   │   └── banner.go
│   │   │   ├── base.go
│   │   │   ├── brands
│   │   │   │   └── brand.go
│   │   │   ├── category
│   │   │   │   └── category.go
│   │   │   └── goods
│   │   │       └── goods.go
│   │   ├── config
│   │   │   └── config.go
│   │   ├── config.yaml
│   │   ├── forms
│   │   │   ├── banner.go
│   │   │   ├── brand.go
│   │   │   ├── category.go
│   │   │   └── goods.go
│   │   ├── global
│   │   │   └── global.go
│   │   ├── initialize
│   │   │   ├── config.go
│   │   │   ├── logger.go
│   │   │   ├── router.go
│   │   │   ├── sentinel.go
│   │   │   └── srv_conn.go
│   │   ├── main.go
│   │   ├── middlewares
│   │   │   ├── admin.go
│   │   │   ├── cors.go
│   │   │   ├── jwt.go
│   │   │   └── tracing.go
│   │   ├── models
│   │   │   └── request.go
│   │   ├── proto
│   │   │   ├── build.sh
│   │   │   ├── gen
│   │   │   │   ├── good.pb.go
│   │   │   │   └── good_grpc.pb.go
│   │   │   └── good.proto
│   │   ├── router
│   │   │   ├── banner.go
│   │   │   ├── brand.go
│   │   │   ├── category.go
│   │   │   └── goods.go
│   │   └── utils
│   │       └── register
│   │           └── consul
│   │               └── register.go
│   ├── order-web
│   │   ├── api
│   │   │   ├── base.go
│   │   │   ├── order
│   │   │   │   └── order.go
│   │   │   ├── pay
│   │   │   │   └── alipay.go
│   │   │   └── shop_cart
│   │   │       └── shop_cart.go
│   │   ├── config
│   │   │   └── config.go
│   │   ├── config.yaml
│   │   ├── forms
│   │   │   ├── order.go
│   │   │   └── shop_cart.go
│   │   ├── global
│   │   │   └── global.go
│   │   ├── initialize
│   │   │   ├── config.go
│   │   │   ├── logger.go
│   │   │   ├── router.go
│   │   │   ├── sentinel.go
│   │   │   └── srv_conn.go
│   │   ├── main.go
│   │   ├── middlewares
│   │   │   ├── cors.go
│   │   │   ├── jwt.go
│   │   │   └── tracing.go
│   │   ├── models
│   │   │   └── request.go
│   │   ├── proto
│   │   │   ├── build.sh
│   │   │   ├── gen
│   │   │   │   ├── goods.pb.go
│   │   │   │   ├── goods_grpc.pb.go
│   │   │   │   ├── inventory.pb.go
│   │   │   │   ├── inventory_grpc.pb.go
│   │   │   │   ├── order.pb.go
│   │   │   │   └── order_grpc.pb.go
│   │   │   ├── goods.proto
│   │   │   ├── inventory.proto
│   │   │   └── order.proto
│   │   ├── router
│   │   │   ├── order.go
│   │   │   └── shop_cart.go
│   │   └── utils
│   │       └── register
│   │           └── consul
│   │               └── register.go
│   ├── oss-web
│   │   ├── config
│   │   │   └── config.go
│   │   ├── config.yaml
│   │   ├── global
│   │   │   └── global.go
│   │   ├── handler
│   │   │   └── oss.go
│   │   ├── initialize
│   │   │   ├── config.go
│   │   │   ├── logger.go
│   │   │   └── router.go
│   │   ├── main.go
│   │   ├── middlewares
│   │   │   └── cors.go
│   │   ├── router
│   │   │   └── oss.go
│   │   ├── static
│   │   │   ├── css
│   │   │   │   └── style.css
│   │   │   └── js
│   │   │       └── upload.js
│   │   ├── templates
│   │   │   └── index.html
│   │   └── utils
│   │       ├── oss.go
│   │       └── register
│   │           └── consul
│   │               └── register.go
│   ├── user-web
│   │   ├── api
│   │   │   ├── chaptcha.go
│   │   │   ├── sms.go
│   │   │   └── user.go
│   │   ├── config
│   │   │   └── config.go
│   │   ├── config.yaml
│   │   ├── forms
│   │   │   ├── sms.go
│   │   │   └── user.go
│   │   ├── global
│   │   │   ├── global.go
│   │   │   └── response
│   │   │       └── user.go
│   │   ├── initialize
│   │   │   ├── config.go
│   │   │   ├── logger.go
│   │   │   ├── router.go
│   │   │   └── srv_conn.go
│   │   ├── main.go
│   │   ├── middleware
│   │   │   ├── admin.go
│   │   │   ├── cors.go
│   │   │   └── jwt.go
│   │   ├── models
│   │   │   └── request.go
│   │   ├── proto
│   │   │   ├── build.sh
│   │   │   ├── gen
│   │   │   │   ├── user.pb.go
│   │   │   │   └── user_grpc.pb.go
│   │   │   └── user.proto
│   │   ├── router
│   │   │   ├── base.go
│   │   │   └── user.go
│   │   ├── utils
│   │   │   └── register
│   │   │       └── consul
│   │   │           └── register.go
│   │   └── validator
│   │       └── validators.go
│   └── userop-web
│       ├── api
│       │   ├── address
│       │   │   └── address.go
│       │   ├── base.go
│       │   ├── message
│       │   │   └── message.go
│       │   └── user_fav
│       │       └── user_fav.go
│       ├── config
│       │   └── config.go
│       ├── config.yaml
│       ├── forms
│       │   ├── address.go
│       │   ├── message.go
│       │   └── user_fav.go
│       ├── global
│       │   └── global.go
│       ├── initialize
│       │   ├── config.go
│       │   ├── logger.go
│       │   ├── router.go
│       │   └── srv_conn.go
│       ├── main.go
│       ├── middlewares
│       │   ├── cors.go
│       │   └── jwt.go
│       ├── models
│       │   └── request.go
│       ├── proto
│       │   ├── address.proto
│       │   ├── build.sh
│       │   ├── gen
│       │   │   ├── address.pb.go
│       │   │   ├── address_grpc.pb.go
│       │   │   ├── good.pb.go
│       │   │   ├── good_grpc.pb.go
│       │   │   ├── message.pb.go
│       │   │   ├── message_grpc.pb.go
│       │   │   ├── userfav.pb.go
│       │   │   └── userfav_grpc.pb.go
│       │   ├── good.proto
│       │   ├── message.proto
│       │   └── userfav.proto
│       ├── router
│       │   ├── address.go
│       │   ├── message.go
│       │   └── user_fav.go
│       └── utils
│           └── register
│               └── consul
│                   └── register.go
├── go.mod
└── srvs
    ├── go.mod
    ├── go.sum
    ├── goods_srv
    │   ├── config
    │   │   └── config.go
    │   ├── config.yaml
    │   ├── global
    │   │   └── global.go
    │   ├── handler
    │   │   ├── banner.go
    │   │   ├── base.go
    │   │   ├── brands.go
    │   │   ├── category.go
    │   │   ├── category_brand.go
    │   │   └── goods.go
    │   ├── initialize
    │   │   ├── config.go
    │   │   ├── db.go
    │   │   ├── es.go
    │   │   └── logger.go
    │   ├── main.go
    │   ├── model
    │   │   ├── base.go
    │   │   ├── es_goods.go
    │   │   ├── goods.go
    │   │   └── main
    │   │       └── main.go
    │   └── proto
    │       ├── build.sh
    │       ├── gen
    │       │   ├── good.pb.go
    │       │   └── good_grpc.pb.go
    │       └── good.proto
    ├── inventory_srv
    │   ├── config
    │   │   └── config.go
    │   ├── config.yaml
    │   ├── global
    │   │   └── global.go
    │   ├── handler
    │   │   └── inventory.go
    │   ├── initialize
    │   │   ├── config.go
    │   │   ├── db.go
    │   │   └── logger.go
    │   ├── main.go
    │   ├── model
    │   │   ├── base.go
    │   │   ├── inventory.go
    │   │   └── main
    │   │       └── main.go
    │   └── proto
    │       ├── build.sh
    │       ├── gen
    │       │   ├── inventory.pb.go
    │       │   └── inventory_grpc.pb.go
    │       └── inventory.proto
    ├── order_srv
    │   ├── config
    │   │   └── config.go
    │   ├── config.yaml
    │   ├── global
    │   │   └── global.go
    │   ├── handler
    │   │   ├── base.go
    │   │   └── order.go
    │   ├── initialize
    │   │   ├── config.go
    │   │   ├── db.go
    │   │   ├── logger.go
    │   │   └── srv_conn.go
    │   ├── main.go
    │   ├── model
    │   │   ├── base.go
    │   │   ├── main
    │   │   │   └── main.go
    │   │   └── order.go
    │   └── proto
    │       ├── build.sh
    │       ├── gen
    │       │   ├── goods.pb.go
    │       │   ├── goods_grpc.pb.go
    │       │   ├── inventory.pb.go
    │       │   ├── inventory_grpc.pb.go
    │       │   ├── order.pb.go
    │       │   └── order_grpc.pb.go
    │       ├── goods.proto
    │       ├── inventory.proto
    │       └── order.proto
    ├── sql
    │   ├── ecs_goods_srv.sql
    │   ├── ecs_inventory_srv.sql
    │   ├── ecs_order_srv.sql
    │   ├── ecs_user_srv.sql
    │   └── ecs_userop_srv.sql
    ├── user_srv
    │   ├── config
    │   │   └── config.go
    │   ├── config.yaml
    │   ├── global
    │   │   └── global.go
    │   ├── handler
    │   │   └── user.go
    │   ├── initialize
    │   │   ├── config.go
    │   │   ├── db.go
    │   │   └── logger.go
    │   ├── main.go
    │   ├── model
    │   │   ├── migrate
    │   │   │   └── main.go
    │   │   └── user.go
    │   ├── proto
    │   │   ├── build.sh
    │   │   ├── gen
    │   │   │   ├── user.pb.go
    │   │   │   └── user_grpc.pb.go
    │   │   └── user.proto
    │   └── utils
    │       └── addr.go
    └── userop_srv
        ├── config
        │   └── config.go
        ├── config.yaml
        ├── global
        │   └── global.go
        ├── handler
        │   ├── address.go
        │   ├── base.go
        │   ├── message.go
        │   └── userfav.go
        ├── initialize
        │   ├── config.go
        │   ├── db.go
        │   └── logger.go
        ├── main.go
        ├── model
        │   ├── base.go
        │   ├── main
        │   │   └── main.go
        │   └── userop.go
        └── proto
            ├── address.proto
            ├── build.sh
            ├── gen
            │   ├── address.pb.go
            │   ├── address_grpc.pb.go
            │   ├── message.pb.go
            │   ├── message_grpc.pb.go
            │   ├── userfav.pb.go
            │   └── userfav_grpc.pb.go
            ├── message.proto
            └── userfav.proto
```

## License

This project is under the Apache License 2.0. See the LICENSE file for the full license text.