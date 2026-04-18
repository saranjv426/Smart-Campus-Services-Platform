// Mock axios for testing
const mockAxios = {
  get: jest.fn(),
  post: jest.fn(),
  put: jest.fn(),
  delete: jest.fn(),
  patch: jest.fn(),
  request: jest.fn(),
  create: jest.fn(),
  interceptors: {
    request: {
      use: jest.fn(),
    },
  },
};

mockAxios.create.mockReturnValue(mockAxios);

export default mockAxios;
