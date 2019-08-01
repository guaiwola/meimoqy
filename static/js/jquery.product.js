jQuery.fn.LoadImage=function(scaling,obj){
	imgLoad(obj);
	var n = obj.length;
	function imgLoad(obj){
		var src = obj.attr("data-img");
		var img = new Image();
		img.src = src;
		//缓存
		if(img.complete || img.width){
			productClass();
		    return;
		}
		$(img).load(function(){
			var newN = obj.length;
			if(newN == n){
				obj.eq(0).show();
			}
			if(obj.next().length){
				imgLoad(obj.next());
			}else{
				productClass();
			}
		});
	}
}
function productClass(){
	$(".proclass > ul > li").each(function(index) {
		$(this).stop().delay(Math.ceil(index)*200).fadeIn(600);
	});
	$(".proclass > ul > li").hover(function(){
		$(this).children(".proclasscover").stop().fadeOut().next().stop().animate({"top":[50,"easeOutCirc"]},600);	
	},function(){
		$(this).children(".proclasscover").stop().fadeIn().next().stop().animate({"top":[60,"easeOutCirc"]},600);
	})
}

$(function(){
	$(".proclass > ul > li").LoadImage(true,$(".proclass > ul > li"));
	
   $(".probox").hover(function(){
		$(this).children(".proboxcover").stop().fadeOut(400);   
	},function(){
		$(this).children(".proboxcover").stop().fadeIn(400);   
	});
	
	$(".prori").find("li").hover(function(){
		$(this).find(".prolicover").stop().fadeIn(100,function(){
			$(this).next().animate({"top":[22,"easeOutCirc"]},400).next().delay(50).animate({"bottom":[18,"easeOutCirc"]},400)
		})	
	},function(){
		$(this).find(".prolicover").stop().fadeOut(350).next().stop().animate({"top":[-240,"easeOutCirc"]},300).next().stop().animate({"bottom":[-40,"easeOutCirc"]},300);
		
	});
	
	$(".probox:odd").css({"background":"#f2f2f2"});
});