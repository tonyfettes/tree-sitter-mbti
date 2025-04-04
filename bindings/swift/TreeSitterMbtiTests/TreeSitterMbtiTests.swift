import XCTest
import SwiftTreeSitter
import TreeSitterMbti

final class TreeSitterMbtiTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_mbti())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading MoonBit Interface grammar")
    }
}
