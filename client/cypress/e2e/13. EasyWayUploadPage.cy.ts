describe('EasyWay Upload Picture Page Test', () => {

  before(() => {
    cy.visit('http://localhost:4200/upload');
  });

  it('should contain a header section with a dark background and a title that says "Upload a Picture"', () => {
    cy.get('.service-header')
      .should('have.css', 'background-color', 'rgb(33, 37, 41)')
      .find('h1')
      .should('have.text', 'Upload a Picture');
  });

  it('should contain a file upload section with a form and a submit button', () => {
    cy.get('.upload')
      .should('exist')
      .find('h2')
      .should('have.text', 'Upload a File');
    cy.get('form')
      .should('exist')
      .find('input[type="file"]')
      .should('exist');
    cy.get('form')
      .should('exist')
      .find('button[type="submit"]')
      .should('exist')
      .should('have.text', 'Upload');
  });
});
