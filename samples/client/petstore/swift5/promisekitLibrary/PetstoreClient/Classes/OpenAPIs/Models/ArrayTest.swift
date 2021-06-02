//
// ArrayTest.swift
//
// Generated by openapi-generator
// https://openapi-generator.tech
//

import Foundation
#if canImport(AnyCodable)
import AnyCodable
#endif

public struct ArrayTest: Codable, Hashable {

    public var arrayOfString: [String]?
    public var arrayArrayOfInteger: [[Int64]]?
    public var arrayArrayOfModel: [[ReadOnlyFirst]]?

    public init(arrayOfString: [String]? = nil, arrayArrayOfInteger: [[Int64]]? = nil, arrayArrayOfModel: [[ReadOnlyFirst]]? = nil) {
        self.arrayOfString = arrayOfString
        self.arrayArrayOfInteger = arrayArrayOfInteger
        self.arrayArrayOfModel = arrayArrayOfModel
    }

    public enum CodingKeys: String, CodingKey, CaseIterable {
        case arrayOfString = "array_of_string"
        case arrayArrayOfInteger = "array_array_of_integer"
        case arrayArrayOfModel = "array_array_of_model"
    }

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {
        var container = encoder.container(keyedBy: CodingKeys.self)
        try container.encodeIfPresent(arrayOfString, forKey: .arrayOfString)
        try container.encodeIfPresent(arrayArrayOfInteger, forKey: .arrayArrayOfInteger)
        try container.encodeIfPresent(arrayArrayOfModel, forKey: .arrayArrayOfModel)
    }
}
