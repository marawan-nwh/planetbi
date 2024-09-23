## Tasks

- Authentication handlers and pages
  - Auth pages should redirect to homepage if user is already signed in
  - Review all auth endpoint
- Homepage (Role/Access (View, Editor, Owner))
  - Create viewset (run the migration first)
  - Navigate to the viewset edit page and show a modal for creating a new view
  - Edit page
  - View page
- https://svelte.dev/docs/accessibility-warnings#a11y-click-events-have-key-events
- Account page (view, edit, delete account)
- Add terms and conditions, and check the link in the signup page
- Email sending needs more testing (could this be because it is in sandbox now?)
- Create tests for important endpoints
- Rate limiter (espectially for important auth endpoints)
- If collections are too many, we gonna need a scrollbar

## Actions on Views

- Publish
- Share (if private)

## Actions on Viewsets

- Publish (select which views to publish)
- Share (if private)
