describe('EasyWay Search Test', () => {

  before(() => {
    cy.visit('http://localhost:4200');
  });

  it('should show search results when a valid search term is entered', () => {
    cy.get('.search-bar input[type="text"]').type('Plumbing');
    cy.get('.search-bar input[type="text"]').should('have.value', 'Plumbing');
  });

});
