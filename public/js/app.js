(function($){
    var busy = false;
    var favorites = [];

    var $removeButton = $('<button class="removeName">&times;</button>');
    //var $favoriteButton = $('<button id="addToFavorites">&#9733;</button>');

    function loadNames() {
        var names = JSON.parse(localStorage.getItem('favoriteNames'));
        if (names.constructor == Array) {
            favorites = names;
        } else {
            favorites = [];
        }
        refreshFavorites();
    }

    function persistFavorites() {
        localStorage.setItem('favoriteNames',JSON.stringify(favorites));
    }

    function addName(name) {
        if (favorites.indexOf(name) >= 0) {
            return false;
        }
        favorites.push(name);
        persistFavorites();
        refreshFavorites();
        return true;
    }

    function removeName(index) {
        favorites.splice(index, 1);
        persistFavorites();
        refreshFavorites();
    }

    function refreshFavorites() {
        favorites.sort();
        $newFavoritesList = $('<ul></ul>');
        for (i=0; i<favorites.length; i++) {
            var $item = $('<li></li>');
            $item.attr('data-index',i).text(favorites[i]);
            $removeButton.clone().appendTo($item);
            $item.appendTo($newFavoritesList);
        }
        $('#favoritesList').html($newFavoritesList.html());
        refreshGeneratedNameAttributes()
    }

    function refreshGeneratedNameAttributes() {
        $currentName = $('#generatedName span').last();
        if (favorites.indexOf($currentName.text()) >= 0) {
            $('.generator').addClass('is-favorite');
        } else {
            $('.generator').removeClass('is-favorite');
        }
    }

    

    $(document).ready(function(){
        loadNames();
        $('#generateName').click(function(e){
            if (busy) {
                return false;
            }
            busy = true;
            $('#generatedName span').animate({
                top: '2em',
                opacity: 0
            }, 200, function(){
                $(this).remove();
            });
            $.get('/generate', function(data){
                $newName = $('<span></span>').css({
                    top: 0,
                    opacity: 0,
                });
                $newName.text(data);
                //$favoriteButton.appendTo($newName);
                $newName.appendTo($('#generatedName'));
                $newName.animate({
                    top: '1em',
                    opacity: 1,
                }, 200);
                refreshGeneratedNameAttributes()
                busy = false;
            },'text');
        });
        $('body').on('click','#addToFavorites',function(){
            nameIndex = favorites.indexOf($('#generatedName span').last().text());
            if ( nameIndex >= 0) {
                return removeName(nameIndex);
            }
            return addName($('#generatedName span').last().text());
        });
        $("body").on("click",".removeName", function(){
            removeName($(this).parents('li').data('index'));
        });
        $('#generateName').trigger('click');
    });

})(jQuery);