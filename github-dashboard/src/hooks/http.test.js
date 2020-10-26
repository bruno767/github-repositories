import React from "react";
import {useHttp} from "./http";
import "whatwg-fetch";
import {renderHook} from "@testing-library/react-hooks";
import fetchMock from "fetch-mock";
import {act} from "react-test-renderer";
const fetch = require('jest-fetch-mock');


describe("http", () => {
    beforeAll(() => {
        jest.setMock('node-fetch', fetch);
    });
    afterAll(() => {
        fetchMock.restore();
    });

    it("should return data with a successful request", async () => {
        const {result} = renderHook(() => useHttp());
        fetchMock.mock("http://api.com/repositories", {
            name: "Joni Baez",
            html_url: "https://some.com",
            description: "some desc"
        });
        await act(async () => {
            result.current.callApi('http://api.com/repositories')
        });

        expect(result.current.loading).toBe(false);
        expect(result.current.data).toMatchObject({
            name: "Joni Baez",
            html_url: "https://some.com",
            description: "some desc"
        });

    });
});