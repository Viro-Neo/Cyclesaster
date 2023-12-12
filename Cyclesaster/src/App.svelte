<script lang="ts">
    import { Router, Route } from 'svelte-routing';
    import LandingPage from './screens/Home.svelte';
    import Map from './screens/Map.svelte';
    import Chart from './screens/Chart.svelte';

    const tableData = [
        { label: 'A', value: 10 },
        { label: 'B', value: 20 },
        { label: 'C', value: 15 },
        { label: 'D', value: 5 },
        { label: 'E', value: 8 },
        { label: 'F', value: 12 },
        { label: 'G', value: 7 },
        { label: 'H', value: 9 },
        { label: 'I', value: 13 },
        { label: 'J', value: 15 },
        { label: 'K', value: 20 },
        { label: 'L', value: 25 },
        { label: 'M', value: 30 },
        { label: 'N', value: 35 },
        { label: 'O', value: 40 },
        { label: 'P', value: 45 },
        { label: 'Q', value: 50 },
        { label: 'R', value: 55 },
        { label: 'S', value: 60 },
        { label: 'T', value: 65 },
        { label: 'U', value: 70 },
        { label: 'V', value: 75 },
        { label: 'W', value: 80 },
        { label: 'X', value: 85 },
        { label: 'Y', value: 90 },
        { label: 'Z', value: 95 },
        // Add more data as needed
    ];
</script>

<Router>
    <Route path="/" component={LandingPage} />
    <Route path="/map" component={Map} />
    <Route path="/chart" let:queryParams>
        <Chart {tableData} />
    </Route>
</Router>
