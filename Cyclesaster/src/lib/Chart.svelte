<script lang="ts">
    import { onMount } from 'svelte';
    import Chart from 'chart.js/auto';

    export let tableData: { label: string; value: number }[];

    let chart: Chart;

    onMount(() => {
        const labels = tableData.map((data) => data.label);
        const values = tableData.map((data) => data.value);

        const ctx = document.getElementById('chart');
        if (ctx) {
            chart = new Chart(ctx, {
                type: 'bar',
                data: {
                    labels,
                    datasets: [
                        {
                            label: 'Chart Data',
                            data: values,
                            backgroundColor: 'rgba(54, 162, 235, 0.5)', // Adjust color as needed
                            borderColor: 'rgba(54, 162, 235, 1)', // Adjust color as needed
                            borderWidth: 1,
                        },
                    ],
                },
                options: {
                    responsive: true,
                    maintainAspectRatio: false,
                    // Add more chart options as needed
                },
            });
        }
    });
</script>

<style>
    /* Add styles as needed */
    #chart-container {
        width: 70lh;
        height: 50vh;
    }
</style>

<div id="chart-container">
    <canvas id="chart"></canvas>
</div>
