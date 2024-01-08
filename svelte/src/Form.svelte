<svelte:options customElement={{
    tag: "property-registration-form",
    shadow:"none", 
}} />
<script>
	import InputField from "./InputField.svelte";

	let steps = [
		'General Information',
		'Location Details',
		'Property Details',
		'Additional Information'
	];
	let currentActive = 0;

    $: activeStep = steps[currentActive];

	function handlePrevious() {
		if (currentActive > 0) currentActive--;
	}

	function handleNext() {
		if (currentActive < steps.length - 1) currentActive++;
	}


</script>

<div class="grid h-full">
	<div class="w-10/12 h-5/6 bg-white rounded-lg place-self-center grid place-content-center">
		<form class="grid grid-cols-3 w-full h-full">
	<section class={activeStep == 'General Information' ? 'block': 'hidden'}>
		<InputField name="Listing Name" />
		<InputField name="Property Type"  />
		<InputField name="Description" />
		<InputField name="Price" />
		<InputField name="Currency"  />
	</section>
	<section class={activeStep == 'Location Details' ? 'block': 'hidden'}>
		<InputField name="Address"  />
		<InputField name="Neighborhood/Area" />
	</section>
	<section class={activeStep == 'Property Details' ? 'block': 'hidden'}>
		<InputField name="Size" />
		<InputField name="Number of Bedrooms"/>
		<InputField name="Number of Bathrooms" />
		<InputField name="Year Built" />
		<InputField name="Flooring Type"/>
		<InputField name="Amenities"/>
		<InputField name="Special Features" />
	</section>
	<section class={activeStep == 'Additional Information' ? 'block': 'hidden'}>
		<InputField name="Availability" />
		<InputField name="Status"  />
	</section>
    <div class="flex flex-row w-full justify-evenly">
			<button on:click={handlePrevious} class="btn btn-primary w-20" type="button" disabled={currentActive === 0}
				>Previous</button
			>
			<button
				on:click={activeStep === "Additional Information" ?  null : handleNext()}
				type={activeStep === "Additional Information" ? "submit": "button"}
				class="btn btn-primary w-20"
				disabled={currentActive === steps.length - 1}>{activeStep === "Additional Information" ? "Submit": "Next" }</button
			>
		</div>
</form>
		
	</div>
</div>
