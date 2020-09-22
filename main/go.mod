module main

go 1.15

replace local/bubblesort => ../bubblesort

replace local/selectionsort => ../selectionsort

replace local/insertionsort => ../insertionsort

require (
	local/bubblesort v0.0.0-00010101000000-000000000000
	local/insertionsort v0.0.0-00010101000000-000000000000
	local/selectionsort v0.0.0-00010101000000-000000000000
)
