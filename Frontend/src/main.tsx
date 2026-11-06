import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import App from './App'

// "!" = we know #root exists (it's in index.html), so it's never null
createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App />
  </StrictMode>,
)
