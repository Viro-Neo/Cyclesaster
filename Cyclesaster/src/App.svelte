<script lang="ts">
    import Map from './lib/Map.svelte';
    import Chart from './lib/Chart.svelte';

    const tableData = [
        { label: 'A', value: 10 },
        { label: 'B', value: 20 },
        { label: 'C', value: 15 },
        // Add more data as needed
    ];
</script>

<div>
    <Chart {tableData} />
    <Map />
</div>
