<script lang="ts">
  import { fade } from "svelte/transition";
  import count from "./store";
</script>

<!-- <div> -->
  <header>
    <p>hi, this counts the number of borpa's in moonmoons chat</p>
  </header>

  <main class="container">
    <span class="count"><code transition:fade>{$count}</code></span>
  </main>
<!-- </div> -->

<style>
  :global(*) {
    box-sizing: border-box;
  }
  :global(body) {
    margin: 0;
    padding: 0;
    font-family: Arial, Helvetica, sans-serif;
    /* height: 100vh; */
    max-height: 100vh;
    max-width: 100vw;
    /* overflow: hidden; */
    position: relative;
  }

  header {
    text-align: center;
    height: 75px;
    display: flex;
    justify-content: center;
    align-items:center;
  }

  header p {
    margin: 0;
  }

  .container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    /* height: 100%; */
  }

  .container code {
    background: #0002;
    padding: 4px 8px;
    border-radius: 4px;
  }

  .count {
    font-size: 3rem;
  }
</style>
