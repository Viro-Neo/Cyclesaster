<script lang="ts">
    import { Router, Route } from 'svelte-routing';
    import LandingPage from './screens/Home.svelte';
    import Map from './screens/Map.svelte';
    import Chart from './screens/Chart.svelte';
</script>

<Router>
    <Route path="/" component={LandingPage} />
    <Route path="/map" component={Map} />
    <Route path="/chart" component={Chart} />
</Router>
