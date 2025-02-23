<template>
  <main>
    <h1>Login</h1>
    <form @submit.prevent="handleLogin">
      <div class="input-group">
        <label for="username">Username</label>
        <input id="username" v-model="username" type="text" required />
      </div>

      <div class="input-group">
        <label for="password">Password</label>
        <input id="password" v-model="password" type="password" required />
      </div>

      <button type="submit">Log In</button>
      <p v-if="error" class="error">{{ error }}</p>
    </form>
  </main>
</template>

<script setup>
try {
  const response = await fetch("https://localhost:8080/api/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ username: username.value, password: password.value }),
  });

  const data = await response.json();

  if (!response.ok) {
    throw new Error(data.message || "Login failed");
  }

  localStorage.setItem("userToken", data.token);

  // todo: goto /chat

} catch (err) {
  console.log(err)
}
</script>

<style scoped>
main {
  max-width: 400px;
  margin: 50px auto;
  padding: 20px;
  border-radius: 10px;
}

h1 {
  text-align: center;
  font-weight: 700;
  font-size: 1.5rem;
}

form {
  display: flex;
  flex-direction: column;
}

.input-group {
  margin-bottom: 15px;
}

label {
  font-size: 0.9rem;
  font-weight: 600;
}

input {
  width: 100%;
  padding: 10px;
  margin-top: 5px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

button {
  background-color: #000;
  color: white;
  padding: 10px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 1rem;
}

button:hover {
  background-color: #333;
}

.error {
  color: red;
  text-align: center;
  margin-top: 10px;
}
</style>
