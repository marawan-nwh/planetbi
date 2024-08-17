<script>
  import { title } from "$lib/stores.js";

  let _title = "";
  title.subscribe((val) => {
    _title = val;
  });
</script>

<svelte:head>
  <title>{_title} - PlanetBI</title>
</svelte:head>

<slot></slot>
