# first-application

## 概要

[DMM-MicroServiceArchitectのオンボーディング課題](https://confl.arms.dmm.com/pages/viewpage.action?pageId=626261798)

## ディレクトリ構成

```bash
.
├── app
│   ├── bff
│   │   ├── handler
│   │   ├── router
│   │   └── util
│   └── server
│       ├── batch
│       ├── handler
│       ├── router
│       └── util
└── eks
    ├── base
    │   ├── bff
    │   └── server
    └── overlays
        └── dev
            ├── bff
            └── server
```

## 技術スタック

- Go(Echo) + AWS(EKS)
