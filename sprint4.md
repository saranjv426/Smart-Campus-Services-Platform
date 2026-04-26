# Sprint 4 Report - Smart Campus Services Platform

A comprehensive Sprint 4 summary for the Smart Campus Services Platform focusing on expanding frontend test coverage for critical user authentication and service discovery flows.

**Status**: ✅ **Sprint 4 In Progress** | **Last Updated**: April 26, 2026 | **Sprint Focus**: Frontend Test Coverage for Authentication, Services Discovery, and Service Details

---

## 🎯 Sprint 4 Overview

### Team Members
- Venkata Sai Saran Jonnalagadda - 11114995
- Srikar Panuganti - 38909216
- Keerthi Reddy Gudibandi - 13652831
- Vishnu Sai Padyala - 32712860

### Sprint Goal
- Achieve comprehensive test coverage for critical frontend user flows including authentication (Login/Register), service discovery (Services Page with filtering and search), and service booking (ServiceDetail Page).

### Sprint 4 Completion Summary

| Area | Outcome | Status |
|------|---------|--------|
| Services Page Filter & Search Tests (#85) | In Progress | 🔄 |
| ServiceDetail Page Tests (#84) | In Progress | 🔄 |
| Login & Register Page Tests (#83) | In Progress | 🔄 |
| End-to-end test scenarios | In Progress | 🔄 |
| Test coverage documentation | In Progress | 🔄 |

---

## 📋 Work Planned For Sprint 4

## 1. Services Page Filter & Search Tests (#85)

### Description
Comprehensive test coverage for the Services page filtering and search functionality, enabling users to discover services by category and search terms.

### Test Coverage
- **Category Filter Tests**:
  - Filtering services by library category
  - Filtering services by dining category
  - Filtering services by transportation category
  - Filtering services by health category
  - Filtering services by housing category
  - Filtering services by events category
  - Filtering services by outdoor category
  - Clearing selected filters

- **Search Functionality Tests**:
  - Searching services by name
  - Searching services by description
  - Partial name matching
  - Case-insensitive search
  - Clearing search input

- **Combined Filter & Search Tests**:
  - Applying filter and search together
  - Clearing all filters and search
  - Filter interaction with search results

- **Edge Cases & State Tests**:
  - Empty services list
  - No results matching search/filter
  - API error handling
  - Loading state display
  - Error state display

### Key Frontend Files
- `frontend/src/pages/Services.js`
- `frontend/src/pages/Services.test.js`

### Test Framework
- Jest for unit testing
- React Testing Library for component testing
- Mock axios for API calls

---

## 2. ServiceDetail Page Tests (#84)

### Description
Comprehensive test coverage for the ServiceDetail page, which displays detailed service information and allows users to submit booking requests.

### Test Coverage
- **Service Detail Rendering Tests**:
  - Service name, description, and location display
  - Service rating and review count display
  - Service image rendering
  - Service availability status

- **Booking Form Tests**:
  - Form field presence validation
  - Date/time picker interaction
  - Number of people validation
  - Booking notes/preferences text area
  - Submit button state management

- **Booking Form Validation Tests**:
  - Required field validation (date, time, duration)
  - Invalid date selection (past dates)
  - Overlapping booking detection
  - User feedback on validation errors
  - Success message after submission

- **Review & Rating Tests**:
  - Existing reviews display
  - Review rating display
  - Review author and date display
  - Review submission form (if applicable)
  - No reviews state

- **Navigation & State Tests**:
  - Back button navigation
  - Service not found error state
  - API error handling
  - Loading state display

### Key Frontend Files
- `frontend/src/pages/ServiceDetail.js`
- `frontend/src/pages/ServiceDetail.test.js`

### Test Framework
- Jest for unit testing
- React Testing Library for component testing
- Mock axios for API calls

---

## 3. Login & Register Page Tests (#83)

### Description
Comprehensive test coverage for authentication flows including login and user registration with role selection.

### Test Coverage

#### Login Page Tests
- **Form Rendering**:
  - Email input field presence
  - Password input field presence
  - Login button presence
  - Register link presence

- **Form Validation**:
  - Empty email validation
  - Invalid email format detection
  - Empty password validation
  - Minimum password length enforcement
  - Real-time validation feedback

- **Authentication Tests**:
  - Successful login with valid credentials
  - Failed login with invalid email
  - Failed login with incorrect password
  - API error handling (500, network errors)
  - Login disabled state during API call

- **Token & Storage Tests**:
  - Auth token storage in localStorage
  - User profile storage in localStorage
  - Token persistence across page reload
  - Logout token removal

- **Navigation Tests**:
  - Redirect to home page after successful login
  - Redirect to login page for unauthorized access
  - Remember me functionality (if applicable)

#### Register Page Tests
- **Form Rendering**:
  - All input fields present (email, password, firstName, lastName, phone)
  - Role selection dropdown/buttons
  - Submit button presence
  - Login link presence

- **Form Validation**:
  - Email format validation
  - Password strength validation
  - Password confirmation matching
  - First name and last name required fields
  - Phone number format validation
  - Role selection required

- **Registration Tests**:
  - Successful registration with all valid inputs
  - Duplicate email error handling (409)
  - Password mismatch detection
  - Required field validation messages
  - API error handling

- **Role Selection Tests**:
  - Student role selection
  - Staff role selection (if applicable)
  - Admin role selection (if applicable)
  - Default role selection
  - Role persistence in user profile

- **Post-Registration Tests**:
  - Auto-login after successful registration
  - Redirect to home page or onboarding
  - User profile created with correct role
  - Auth token generation and storage

### Key Frontend Files
- `frontend/src/pages/Login.js`
- `frontend/src/pages/Login.test.js`
- `frontend/src/pages/Register.js`
- `frontend/src/pages/Register.test.js`

### Test Framework
- Jest for unit testing
- React Testing Library for component testing
- User Event for simulating user interactions
- Mock axios for API calls
- Mock react-router-dom for navigation testing

---

## 📊 Test Execution Guidelines

### Running Tests
```bash
cd frontend
npm test
```

### Test File Organization
- `*.test.js` files co-located with component files
- Setup file: `src/setupTests.js`
- Mock utilities: `src/__mocks__/axios.js`
- Testing utilities: `src/utils/testUtils.js`

### Coverage Goals
- Minimum 80% statement coverage for critical pages
- All user interaction paths covered
- All API integration points tested
- Error states and edge cases included

### Test Data
- Mock user data for authentication tests
- Mock service data for discovery tests
- Mock API responses for integration testing
- Sample dates and times for booking tests

---

## 🔗 Related Test Files

| Feature | Test File |
|---------|-----------|
| Services Page | `frontend/src/pages/Services.test.js` |
| ServiceDetail Page | `frontend/src/pages/ServiceDetail.test.js` |
| Login Page | `frontend/src/pages/Login.test.js` |
| Register Page | `frontend/src/pages/Register.test.js` |
| App Component | `frontend/src/App.test.js` |

---

## ✅ Verification Steps

### Local Testing
1. Run `npm test` in frontend directory
2. Verify all tests pass
3. Check coverage reports for target pages
4. Run Cypress e2e tests (if applicable)

### Code Quality
1. Verify test files follow Jest conventions
2. Check mock data accuracy
3. Review test descriptions for clarity
4. Ensure no hardcoded values in tests

### Git Integration
1. All test files committed to repository
2. Tests run successfully in CI/CD pipeline
3. No console errors or warnings in test output
4. Coverage thresholds met

---

## 📝 Notes

- All tests use React Testing Library best practices (querying by role/label, not implementation details)
- Mock axios responses should match actual API contract
- Tests should be independent and not rely on execution order
- Use `beforeEach` and `afterEach` for setup and cleanup
- Test descriptions should clearly state what is being tested and expected outcome
