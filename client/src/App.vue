<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <header>
    <div class="container">
      <div class="left-section">
        <div class="logo-contain">
          <img alt="login" class="logo" src="@/assets/logo.png" width="30" height="30" />
          <span>yaca</span>
        </div>
        <nav>
          <RouterLink to="/" class="nav-link">Home</RouterLink>
          <RouterLink to="/about" class="nav-link">About</RouterLink>
          <RouterLink to="/contact" class="nav-link">Contact</RouterLink>
          <RouterLink to="/blog" class="nav-link">Blog</RouterLink>
        </nav>
      </div>
      <div class="right-section">
        <RouterLink to="/login" class="action-link">Login</RouterLink>
        <RouterLink to="/signup" class="action-link second">Sign up</RouterLink>
      </div>
    </div>
  </header>
  <RouterView />
</template>

<style scoped>
header {
  background-color: #fff;
  padding: 1rem 2rem;
}

.container {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.left-section {
  display: flex;
  align-items: center;
}

.logo-contain {
  display: flex;
  align-items: center;
  gap: 8px;
}

.logo {
  display: block;
  width: 30px;
  height: 30px;
}

span {
  font-weight: bold;
  font-size: 1.2rem;
  margin-left: -0.5rem;
  margin-right: 1.4rem;
}

span {
  font-weight: 900;
  color: #000;
}

nav a.nav-link {
  margin-right: 1rem;
  text-decoration: none;
  color: var(--vt-c-text-light-2);
  font-weight: 600;
}

nav a.nav-link.router-link-exact-active {
  color: #111;
}

.right-section {
  display: flex;
}

.right-section .action-link {
  margin-left: 1rem;
  text-decoration: none;
  font-weight: 600;
  color: white;
  background-color: rgb(0, 168, 73);
  padding: 0.35rem 1.95rem;
  border-radius: 999px;
  transition: background 0.3s ease;
}

.right-section > .action-link + .second {
  background-color: #fff;
  border: 1px solid rgb(0, 168, 73);
  color: rgb(0, 168, 73);
}

.right-section .action-link:hover {
  background-color: darkgreen;
}

@media (max-width: 768px) {
  .container {
    flex-direction: row;
  }

  nav {
    display: flex;
  }

  .right-section {
    display: flex;
  }
}
</style>
