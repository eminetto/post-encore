import { describe, expect, test } from "vitest";
import { greeting } from "./foo";

describe("greeting", () => {
    test("should combine string with parameter value", async () => {
        // Call the greeting endpoint like an ordinary function
        const resp = await greeting({ name: "world" });

        expect(resp.greeting).toBe("Hello world!");
    });
});