<script>
  import { title } from "$lib/js/stores.js";

  let _title = "";
  title.subscribe((val) => {
    _title = val;
  });
</script>

<svelte:head>
  <title>{_title === "" ? "Plainbyte" : _title + " | Plainbyte"}</title>
</svelte:head>

<slot></slot>
