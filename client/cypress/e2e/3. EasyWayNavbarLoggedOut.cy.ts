describe('EasyWay Logged Out Navbar Test', () => {
    before(() => {
      cy.visit('http://localhost:4200/');
    });

    it('has the correct title', () => {
      cy.title().should('equal', 'EasyWay');
    });

    it('displays the correct title', () => {
        cy.get('header')
        .should('have.id','header')
        .find('h3')
        .find('a')
        .should('have.attr','routerLink','/home')
        .should('have.text','EasyWay')
    });

    it('has the nav element', () => {
      cy.get('#header')
      .find('nav')
      .should('have.id','nav')
    });
})
