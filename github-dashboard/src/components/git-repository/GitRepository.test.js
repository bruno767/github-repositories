import React from 'react';
import { render } from '@testing-library/react';
import GitRepository from './GitRepository';
import {test} from "@jest/globals";

test('should show GitRepository fields', () => {
    const { getByText } = render(<GitRepository name="facebook" description="some desc" html_url="https://git.com" language="go"/>);
    const name = getByText(/facebook/i);
    const language = getByText(/go/i);
    const description = getByText(/some desc/i);
    expect(name).toBeInTheDocument();
    expect(language).toBeInTheDocument();
    expect(description).toBeInTheDocument();
});
