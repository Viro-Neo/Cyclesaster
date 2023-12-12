<!-- Chart.svelte -->
<script lang="ts">
    import { onMount } from 'svelte';
    import Chart from 'chart.js/auto';
    import { Link } from 'svelte-routing';

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
    /* Set a specific height and width for the chart container */
    #chart-container {
        width: 80%; /* Adjust the width as needed */
        height: 400px; /* Adjust the height as needed */
        margin: 0 auto; /* Center the container horizontally */
    }

    /* Ensure the canvas takes up the full size of its container */
    canvas {
        width: 100%;
        height: 100%;
    }
</style>

<div id="chart-container">
    <canvas id="chart"></canvas>
    <Link to="/map">Go to Map</Link>
</div>
