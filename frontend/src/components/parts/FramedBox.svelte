<script lang="ts">
  export let title = 'Untitled';
</script>

<div class="relative mt-8 mr-0 mb-8 ml-0 p-4 border-2 border-gray-700 rounded-lg">
    <div class="absolute -top-3 right-0 left-0">
        <span class="pr-2 pl-2 bg-white">{title}</span>
    </div>
    <div>
        <slot/>
    </div>
</div>
