## Tasks

- Authentication handlers and pages
  - Auth pages should redirect to homepage if user is already signed in
  - Review all auth endpoint
- Homepage (Role/Access (View, Editor, Owner))
- Account page (view, edit, delete account)
- Create view and viewlist
- Add terms and conditions, and check the link in the signup page
- Email sending needs more testing (could this be because it is in sandbox now?)
- Create tests for important endpoints
- Rate limiter (espectially for important auth endpoints)

## Ideas

- Dependencies
  Let the user add URLs of JS libs to views, viewsets or dashboards, and we inject them into the page so they can be used in the views.
  This will enable the user to do a lot of things that we didn't add or won't add.
