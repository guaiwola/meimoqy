// jQuery.fn.LoadImage=function(a,c){b(c);var d=c.length;function b(g){var h=g.attr("data-img");var f=g.attr("data-url");var e=new Image();e.src=h;$(e).load(function(){g.attr("style","background:url("+h+") no-repeat center top");g.html('<a href="'+f+'" target="_blank"></a>');var j=g.length;if(j==d){g.eq(0).fadeIn(500)}if(g.next().length){b(g.next())}else{picSwitch("#banner","#number",5,true)}})}};function picSwitch(h,f,b,d){var h=$(h);var c=$(f);var k=b;h.css("background","none");var e=h.find("li").size();for(i=0;i<e;i++){c.append("<span></span>")}c.children("span").eq(0).addClass("on");h.find("li").eq(0).fadeIn(400);var j=1;c.children("span").mouseover(function(){j=c.children("span").index(this);a(j)});function a(l){h.find("li").eq(l).stop().fadeIn(400).siblings("li").fadeOut(400);c.children("span").eq(l).addClass("on").siblings("span").removeClass("on")}if(d==true){h.hover(function(){if(g){clearInterval(g)}},function(){g=setInterval(function(){a(j);j++;if(j==e){j=0}},k*500)});var g=setInterval(function(){a(j);j++;if(j==e){j=0}},k*500)}if(d==false){h.hover(function(){g=setInterval(function(){a(j);j++;if(j==e){j=0}},k*500)},function(){if(g){clearInterval(g)}})}}$(function(){$(".index-banner > ul > li").LoadImage(true,$(".index-banner > ul > li"));$(".index-probox").eq(0).fadeIn(function(){picSwitch("#proser01","#pronum01",2,false)});$(".index-probox").eq(0).delay(0).fadeIn(function(){picSwitch("#proser02","#pronum02",2,false)});$(".index-probox").eq(0).delay(0).fadeIn(function(){picSwitch("#proser03","#pronum03",2,false)});$(".index-probox").eq(0).delay(0).fadeIn(function(){picSwitch("#proser04","#pronum04",2,false)})});

jQuery.fn.LoadImage = function(a, c) {
    b(c);
    var d = c.length;
    function b(g) {
        var h = g.attr("data-img");
        var f = g.attr("data-url");
        var e = new Image();
        e.src = h;
        $(e).load(function() {
            g.attr("style", "background:url(" + h + ") no-repeat center top");
            g.html('<a href="' + f + '" target="_blank"></a>');
            var j = g.length;
            if (j == d) {
                g.eq(0).fadeIn(500)
            }
            if (g.next().length) {
                b(g.next())
            } else {
                picSwitch("#banner", "#number", 5, true)
            }
        })
    }
};
function picSwitch(h, f, b, d) {
    var h = $(h);
    var c = $(f);
    var k = b;
    h.css("background", "none");
    var e = h.find("li").size();
    for (i = 0; i < e; i++) {
        c.append("<span></span>")
    }
    c.children("span").eq(0).addClass("on");
    h.find("li").eq(0).fadeIn(400);
    var j = 1;
    c.children("span").mouseover(function() {
        j = c.children("span").index(this);
        a(j)
    });
    function a(l) {
        h.find("li").eq(l).stop().fadeIn(400).siblings("li").fadeOut(400);
        c.children("span").eq(l).addClass("on").siblings("span").removeClass("on")
    }
    if (d == true) {
        h.hover(function() {
            if (g) {
                clearInterval(g)
            }
        },
        function() {
            g = setInterval(function() {
                a(j);
                j++;
                if (j == e) {
                    j = 0
                }
            },
            k * 500)
        });
        var g = setInterval(function() {
            a(j);
            j++;
            if (j == e) {
                j = 0
            }
        },
        k * 500)
    }
    if (d == false) {
        h.hover(function() {
            g = setInterval(function() {
                a(j);
                j++;
                if (j == e) {
                    j = 0
                }
            },
            k * 500)
        },
        function() {
            if (g) {
                clearInterval(g)
            }
        })
    }
}
$(function() {
    $(".index-banner > ul > li").LoadImage(true, $(".index-banner > ul > li"));
    $(".index-probox").eq(0).fadeIn(function() {
        picSwitch("#proser01", "#pronum01", 2, false)
    });
    $(".index-probox").eq(0).delay(0).fadeIn(function() {
        picSwitch("#proser02", "#pronum02", 2, false)
    });
    $(".index-probox").eq(0).delay(0).fadeIn(function() {
        picSwitch("#proser03", "#pronum03", 2, false)
    });
    $(".index-probox").eq(0).delay(0).fadeIn(function() {
        picSwitch("#proser04", "#pronum04", 2, false)
    })
});