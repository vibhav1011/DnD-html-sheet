
$(function () {
	$("#page1").saveMyForm(
		{
			loadInputs: true,
		}
	);
});

$ (function() {
	$("#ins1").on("click",function() {  
		//$(this).css('background-color', 'black');
		//console.log('!works!');
		$(this).toggleClass("active");
		});
});
