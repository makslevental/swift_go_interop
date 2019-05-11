// swift-tools-version:4.2

import PackageDescription

let package = Package(
        name: "Hello",
        products: [
            .library(name: "Hello", targets: ["Hello"]),
        ],
        targets: [
            .systemLibrary(name: "Hello", path: "."),
        ]
)
