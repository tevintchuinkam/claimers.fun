<script lang="ts">
</script>

<section>
	<div
		class="flex flex-col items-center justify-center border-2 border-gray-300 rounded-md p-4 w-full min-h-screen"
	>
		<div class="flex flex-row items-center justify-center w-full h-full">
			<form
				action="/start-game"
				class="flex flex-row items-center space-x-2 justify-center"
			>
				<input type="text" name="name" placeholder="Enter your name" />
				<button
					type="submit"
					class="bg-blue-500 text-white p-2 rounded-md">Play</button
				>
			</form>
		</div>
	</div>
</section>
