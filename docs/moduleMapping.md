# Module mapping bookWurms

Disclaimer: The list is incomplete, but gives a good overview of relevant modules. Especially some accessibility and DevOps modules can just be added on top, if we need more points.

## Safe picks
| Module index | Module description | bookWurms feature | Points |
|--------------|--------------------|-------------------|--------|
| 1 Web |  Major: Use a framework for both the frontend and backend. Use a frontend framework (React, Vue, Angular, Svelte, etc.). Use a backend framework (Express, NestJS, Django, Flask, Ruby on Rails, etc.). Full-stack frameworks (Next.js, Nuxt.js, SvelteKit) count as both if you use both their frontend and backend capabilities. | Svelte and Gin | 2 |
| 1 Web | Major: Implement real-time features using WebSockets or similar technology. Real-time updates across clients. Handle connection/disconnection gracefully. Efficient message broadcasting. | The feed with live updates to posts | 2 |
| 1 Web | Major: Allow users to interact with other users. The minimum requirements are: A basic chat system (send/receive messages between users). A profile system (view user information). A friends system (add/remove friends, see friends list). | Chat, Profile, Friends | 2 |
| 1 Web | Minor: Use an ORM for the database. | Gorm | 1 |
| 1 Web | Minor: Progressive Web App (PWA) with offline support and installability. |gives the "app feeling" instead of website | 1 |
| 1 Web | Minor: Implement advanced search functionality with filters, sorting, and pagination. | feed filters and book search functionality | 1 |
| 1 Web | Minor: File upload and management system. Support multiple file types (images, documents, etc.). Client-side and server-side validation (type, size, format). Secure file storage with proper access control. File preview functionality where applicable. Progress indicators for uploads. Ability to delete uploaded files. | upload ebooks, PDFs  and profile pictures | 1 |
| 3 User Management | Major: Standard user management and authentication. Users can update their profile information. Users can upload an avatar (with a default avatar if none provided). Users can add other users as friends and see their online status. Users have a profile page displaying their information. | Profile page functionality | 2 |
| 3 User Management | Major: Advanced permissions system: View, edit, and delete users (CRUD). Roles management (admin, user, guest, moderator, etc.). Different views and actions based on user role. | maybe good to have moderator / admins | 2 |
| 10 Modules of choice | Major (or minor) BUT we need to justify it | e-book reader | 2 (or 1)
| SUM | | | 16 |

## Potential picks
| Module index | Module description | bookWurms feature | Points |
|--------------|--------------------|-------------------|--------|
| 1 Web | Minor: A complete notification system for all creation, update, and deletion actions. | tbd | 1 |
| 2 Accessibility and Internationalization | Minor: Support for multiple languages (at least 3 languages). Implement i18n (internationalization) system. At least 3 complete language translations. Language switcher in the UI. All user-facing text must be translatable. | English, German, (Spanish?) | 1 
| 3 User Management | Minor: Implement remote authentication with OAuth 2.0 (Google, GitHub, 42, etc.). | tbd | 1 |
| 3 User Management | Minor: Implement a complete 2FA (Two-Factor Authentication) system for the users. | tbd | 1 |
| 3 User Management | Minor: User activity analytics and insights dashboard. | tbd | 1 |
| 7 Devops | Major: Monitoring system with Prometheus and Grafana. Set up Prometheus to collect metrics. Configure exporters and integrations. Create custom Grafana dashboards. Set up alerting rules. Secure access to Grafana. | metrics and dashboards | 2 |
| 7 Devops| Minor: Health check and status page system with automated backups and disaster recovery procedures. | more devops stuff | 1 |


### Comparison to eval sheet

From Chapter V.3 Social and Collaborative Projects:  

These projects emphasize user interaction and community building:
- Social Network: User profiles, posts, comments, likes, friends, real-time chat, and
notifications.
  - Suggested modules: User interaction, Real-time features, Notification system,
Advanced chat, File upload
  - Point potential: 14+ points