describe('EasyWay Footer Test', () => {
  beforeEach(() => {
      cy.visit('/');
    });

    it('should contain a Services section with links', () => {
      cy.get('.footer-dark #footer')
        .contains('Services')
        .siblings('ul')
        .find('li')
        .should('have.length', 5)
        .each(($li) => {
          cy.wrap($li)
            .find('a')
            .should('have.attr', 'href')
        });
    });

    it('should contain an About section with links', () => {
      cy.get('.footer-dark #footer')
        .contains('About')
        .siblings('ul')
        .find('li')
        .not('[href="#"]') // exclude links with href="#"
        .each(($li) => {
          cy.wrap($li)
            .find('a')
            .should('have.attr', 'href')
        });
    });

    it('should have social media icons with links', () => {
      cy.get('.footer-dark #footer')
        .find('.social')
        .find('a')
        .should('have.length', 4)
        .each(($a) => {
          cy.wrap($a)
            .should('have.attr', 'href')
            .and('include', '#');
        });
    });

    it('should have a copyright notice', () => {
      cy.get('.footer-dark #footer')
        .find('.text-muted')
        .should('have.text', '© EasyWay @ Fall 2022 - Spring 2023');
    });
  });
