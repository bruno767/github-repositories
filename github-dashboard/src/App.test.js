import React from 'react';
import App from './App';
import { render } from '@testing-library/react';
import {test} from "@jest/globals";

test('should render App', () => {
  const { getByText } = render(<App/>);
  const title = getByText(/Git repositories/i);
  expect(title).toBeInTheDocument();
});