describe('EasyWay My Blog Page Test', () => {

  before(() => {
    cy.visit('http://localhost:4200/blog');
  });

  it('should contain a header section with a dark background and a title that says "Blog"', () => {
    cy.get('.service-header')
      .should('have.css', 'background-color', 'rgb(33, 37, 41)')
      .find('h1')
      .should('have.text', 'Blog');
  });

  it('should contain a blog section with a list of posts', () => {
    cy.get('.py-5 article')
      .should('exist')
      .find('li.media')
      .should('have.length.gt', 0);
  });

  it('should contain posts with an image, a title, a creation date, a content section, and a "Read More" button', () => {
    cy.get('.py-5 article')
      .find('li.media')
      .each(($post) => {
        cy.wrap($post)
          .should('exist')
          .find('img')
          .should('have.attr', 'src');
        cy.wrap($post)
          .should('exist')
          .find('a[href^="/blog/"]')
          .should('exist')
          .find('h3')
          .should('exist');
        cy.wrap($post)
          .should('exist')
          .find('small.text-muted')
          .should('exist');
        cy.wrap($post)
          .should('exist')
          .find('p')
          .should('exist');
        cy.wrap($post)
          .should('exist')
          .find('a.btn-primary')
          .should('exist');
      });
  });
});
