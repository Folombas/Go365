// Плавная прокрутка к форме
document.addEventListener('DOMContentLoaded', function() {
    // Плавная прокрутка при клике на ссылку "Новый пост"
    const newPostLink = document.querySelector('a[href="#new-post"]');
    if (newPostLink) {
        newPostLink.addEventListener('click', function(e) {
            e.preventDefault();
            document.querySelector('#new-post-form').scrollIntoView({ behavior: 'smooth' });
        });
    }

    // Подтверждение отправки формы (опционально)
    const postForm = document.querySelector('#new-post-form form');
    if (postForm) {
        postForm.addEventListener('submit', function(e) {
            if (!confirm('Опубликовать пост? (+10 EXP)')) {
                e.preventDefault();
            }
        });
    }

    // Анимация появления карточек
    const cards = document.querySelectorAll('.post-card');
    cards.forEach((card, index) => {
        card.style.animation = `fadeIn 0.5s ease forwards ${index * 0.1}s`;
    });
});