<script lang="ts">
  import { GetAllTasks } from '../../wailsjs/go/main/App';
  import type { entities } from '../../wailsjs/go/models';
  import FramedBox from './parts/FramedBox.svelte';

  let tasks: entities.Task[] = [];

  (async () => tasks = await GetAllTasks())();
</script>

<div class="w-full">
    <FramedBox title="Today's Tasks">
        {#each tasks as task}
            <div class="w-full flex flex-row">
                <div class="mr-2">{task.name}</div>
                <div>done / in progress / not started</div>
            </div>
        {/each}
    </FramedBox>
</div>
