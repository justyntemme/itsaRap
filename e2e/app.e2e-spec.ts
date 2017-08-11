import { ItsaRapPage } from './app.po';

describe('itsa-rap App', () => {
  let page: ItsaRapPage;

  beforeEach(() => {
    page = new ItsaRapPage();
  });

  it('should display welcome message', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('Welcome to app!');
  });
});
