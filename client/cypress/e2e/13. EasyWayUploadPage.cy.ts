describe('EasyWay Upload Picture Page Test', () => {
  beforeEach(() => {
    cy.visit('http://localhost:4200/upload');
  });

  it('should display the header with the correct title', () => {
    cy.get('.service-header')
      .should('be.visible')
      .find('h1')
      .should('have.class', 'display-4')
      .should('have.text', 'Upload a Picture');
  });
});
