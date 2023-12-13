<!-- Chart.svelte -->
<script lang="ts">
    import { onMount } from 'svelte';
    import Chart from 'chart.js/auto';
    import { Link } from 'svelte-routing';

    let selectedFilter = ['all', 'A', 'B', 'C', 'D', 'E'];

    const tableData = [
        { name: 'A', value: 10 },
        { name: 'B', value: 20 },
        { name: 'C', value: 30 },
        { name: 'D', value: 40 },
        { name: 'E', value: 50 },
        // Add more data as needed
    ];

    let chart: Chart;
    const labels = tableData.map((d) => d.name);
    const values = tableData.map((d) => d.value);
    console.log(labels, values);

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
                        }
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

    /* Add this CSS to position the filters */
    .filter-container {
        display: flex;
        align-items: flex-start;
        width: 50%;
        margin: 0 auto;
        border: 0 solid black;
        flex-wrap: wrap; /* Allow items to wrap to the next line */
    }

    .filter {
        width: 10%; /* Adjust the width as needed */
        border: 0 solid black;
        margin: 5px 0; /* Adjust margin for spacing */
    }

    .filter-left {
        text-align: left;
        border: 0 solid black;
    }

    .filter-right {
        text-align: right;
        width: 80%; /* Adjust width for the right filter */
    }

</style>

<!-- Chart.svelte -->
<div id="chart-container">
    <canvas id="chart"></canvas>
    <Link to="/map">Go to Map</Link>
</div>

<div class="filter-container">
    <div class="filter filter-left">
        <select bind:value={selectedFilter}>
            <option value="all">All</option>
            <option value="A">A</option>
            <option value="B">B</option>
            <option value="C">C</option>
            <option value="D">D</option>
            <option value="E">E</option>
            <!-- Add more filter options as needed -->
        </select>
    </div>

    <div class="filter filter-left">
        <select bind:value={selectedFilter}>
            <option value="all">All</option>
            <option value="A">A</option>
            <option value="B">B</option>
            <option value="C">C</option>
            <option value="D">D</option>
            <option value="E">E</option>
            <!-- Add more filter options as needed -->
        </select>
    </div>

    <div class="filter filter-right">
        <select bind:value={selectedFilter}>
            <option value="all">All</option>
            <option value="A">A</option>
            <option value="B">B</option>
            <option value="C">C</option>
            <option value="D">D</option>
            <option value="E">E</option>
            <!-- Add more filter options as needed -->
        </select>
    </div>
</div>
