<!-- Chart.svelte -->
<script lang="ts">
    import { onMount } from 'svelte';
    import Chart from 'chart.js/auto';
    import { Link } from 'svelte-routing';

    const tableData = [
        { year: '2019', month: 'Jan', value: 10 },
        { year: '2019', month: 'Feb', value: 20 },
        { year: '2019', month: 'Mar', value: 30 },
        { year: '2019', month: 'Apr', value: 40 },
        { year: '2019', month: 'May', value: 50 },
        { year: '2019', month: 'Jun', value: 60 },
        { year: '2019', month: 'Jul', value: 70 },
        { year: '2019', month: 'Aug', value: 80 },
        { year: '2019', month: 'Sep', value: 90 },
        { year: '2019', month: 'Oct', value: 100 },
        { year: '2019', month: 'Nov', value: 110 },
        { year: '2019', month: 'Dec', value: 120 },
        { year: '2020', month: 'Jan', value: 130 },
        { year: '2020', month: 'Feb', value: 140 },
        { year: '2020', month: 'Mar', value: 150 },
        { year: '2020', month: 'Apr', value: 160 },
        { year: '2020', month: 'May', value: 170 },
        { year: '2020', month: 'Jun', value: 180 },
        { year: '2020', month: 'Jul', value: 190 },
        { year: '2020', month: 'Aug', value: 200 },
        { year: '2020', month: 'Sep', value: 210 },
        { year: '2020', month: 'Oct', value: 220 },
        { year: '2020', month: 'Nov', value: 230 },
        { year: '2020', month: 'Dec', value: 240 },
        { year: '2021', month: 'Jan', value: 250 },
        { year: '2021', month: 'Feb', value: 260 },
        { year: '2021', month: 'Mar', value: 270 },
        { year: '2021', month: 'Apr', value: 280 }
        // Add more data as needed
    ];

    let chart: Chart;
    let firstFilter = '';
    let secondFilter = '';
    let labels: string[] = [];
    let values: number[] = [];

    onMount(() => {
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
        height: 30lh; /* Adjust the height as needed */
        margin: 0 auto; /* Center the container horizontally */
    }

    /* Ensure the canvas takes up the full size of its container */
    canvas {
        width: 100%;
        height: 100%;
    }
</style>

<div id="chart-container">
    <div>
        <label for="first-filter">First Filter:</label>
        <select id="first-filter" bind:value={firstFilter}>
            <option value="all">All</option>
            <option value="A">A</option>
            <option value="B">B</option>
            <!-- Add more options as needed -->
        </select>
    </div>
    <div>
        <label for="second-filter">Second Filter:</label>
        <select id="second-filter" bind:value={secondFilter}>
            <option value="all">All</option>
            <option value="10">10</option>
            <option value="20">20</option>
            <!-- Add more options as needed -->
        </select>
    </div>

    <canvas id="chart"></canvas>
    <Link to="/map">Go to Map</Link>
</div>
