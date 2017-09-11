(function() {
    var data = {
        id: 3,
        title: 'Mastering Kubernetes',
        img: 'https://images-na.ssl-images-amazon.com/images/I/41deaLZesLL._SX404_BO1,204,203,200_.jpg',
        categories: [{
            name: 'Software Development'
        }],
        description: 'This practical guide demystifies Kubernetes and ensures that your clusters are always ' +
        'available, scalable, and up to date',
        published: 'May 25, 2017'
    };

    var bookBtn = document.getElementById('add-book');

    bookBtn.addEventListener('click', function() {
       fetch('/api/books', {
           method: 'post',
           body: JSON.stringify(data)
       });
    });
})();