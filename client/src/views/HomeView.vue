<template>
  <main>
    <h1>Stay Connected, Anytime, Anywhere</h1>
    <p class="description">A fast and secure way to chat—just you and one other person. No distractions, no clutter.</p>
    <div class="preview">
      <div class="preview-block">
        <p>VIDEO PREVIEW</p>
      </div>
    </div>
  </main>
</template>

<style scoped>
h1 {
  font-weight: 900;
  text-align: center;
  font-size: 2.5rem;
  margin-top: 4rem;
  color: #111111;
}

.description {
  text-align: center;
  color: #555
}

button:hover {
  background-color: #ec5353;
}

.preview {
  margin-top: 4rem;
  display: flex;
  justify-content: center;
}

.preview > .preview-block {
  background-color: #151515;
  border-radius: 12px;
  width: 1000px;
  height: 510px;
  display: flex;
  justify-content: center;
  place-items: center;
}

.preview > .preview-block > p {
  background-color: #fff;
  border-radius: 999px;
  padding: 10px 26px;
  font-weight: 700;
}
</style>
