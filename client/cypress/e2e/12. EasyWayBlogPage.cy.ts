describe('EasyWay My Blog Page Test', () => {
  before(() => {
    cy.visit('http://localhost:4200/blog');
  });

  it('should display the blog header', () => {
    cy.get('header.bg-dark') // Locate the header element
      .should('be.visible') // Ensure it's visible
      .within(() => {
        cy.get('h1.display-4') // Locate the heading element
          .should('have.text', 'Blog'); // Ensure the heading text is 'Blog'
      });
  });

  it('should display at least one blog post', () => {
    cy.get('section.py-5') // Locate the blog section
      .should('be.visible') // Ensure it's visible
      .within(() => {
        cy.get('li') // Locate the blog post elements
          .should('have.length.greaterThan', 0); // Ensure there is at least one post
      });
  });

  it('should load comments for a blog post when the "Load comments" button is clicked', () => {
    cy.get('section.py-5') // Locate the blog section
      .should('be.visible') // Ensure it's visible
      .within(() => {
        cy.get('li') // Locate the blog post elements
          .first() // Select the first post
          .within(() => {
            cy.get('button') // Locate the "Load comments" button
              .should('be.visible') // Ensure it's visible
              .click(); // Click the button
            cy.get('ul') // Locate the comment section
              .should('be.visible'); // Ensure it's visible
          });
      });
  });

  it('should load comments when "Load comments" button is clicked', () => {
     cy.get('button') // Locate all the "Load comments" buttons
       .should('be.visible') // Ensure they're visible
       .each(($button) => { // Iterate through each button
         cy.wrap($button) // Wrap the button element so we can chain Cypress commands
           .click(); // Click the button
         cy.get('form') // Locate the comment form for this button
           .should('be.visible') // Ensure it's visible
           .within(() => {
             // ... add comment to the form ...
           });
       });
   });

});
