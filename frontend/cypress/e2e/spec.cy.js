describe('Basic Test', () => {
  it('visits app and navigates to register page', () => {
    cy.visit('http://localhost:3000');

    cy.contains(/get started/i).click();

    // ✅ Assert navigation
    cy.url().should('include', '/register');

    // ✅ Assert page content
    cy.contains('Register');
  });
});