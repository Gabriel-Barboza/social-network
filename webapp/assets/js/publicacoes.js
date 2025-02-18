$('#nova-publicacao').on('submit', criarpublicacao); 
$(document).on('click', '.curtir-publicacao', curtirpublicacao);
$(document).on('click', '.descurtir-publicacao', descurtirpublicacao);
$('#atualizar-publicacao').on('click', atualizarPublicacao);

function criarpublicacao(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function() {
        window.location.href = "/home";

    }).fail(function(erro) {
        console.log(erro);
        alert("Erro ao cadastrar publicação");
    });

}

function curtirpublicacao(evento) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoid = elementoClicado.closest('div').data('publicacao-id');
    elementoClicado.prop('disabled', true);
    $.ajax({
        url: `/publicacoes/${publicacaoid}/curtir`,
        method: "POST"
    }).done(function() {
        const contadorCurtidas = elementoClicado.next('span');
        const quantidadeAtual = parseInt(contadorCurtidas.text());
        contadorCurtidas.text(quantidadeAtual + 1);

        elementoClicado.addClass('descurtir-publicacao');
        elementoClicado.addClass('text-danger');
        elementoClicado.removeClass('curtir-publicacao');

    }).fail(function(erro) {
        console.log(erro);
        alert("Erro ao curtir publicação");
    }).always(function() {
        elementoClicado.prop('disabled', false);
    });

}

function descurtirpublicacao(evento) {
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoid = elementoClicado.closest('div').data('publicacao-id');
    elementoClicado.prop('disabled', true);
    $.ajax({
        url: `/publicacoes/${publicacaoid}/descurtir`,
        method: "POST"
    }).done(function() {
        const contadorCurtidas = elementoClicado.next('span');
        const quantidadeAtual = parseInt(contadorCurtidas.text());
        contadorCurtidas.text(quantidadeAtual - 1);

        elementoClicado.removeClass('descurtir-publicacao');
        elementoClicado.removeClass('text-danger');
        elementoClicado.addClass('curtir-publicacao');

    }).fail(function(erro) {
        console.log(erro);
        alert("Erro ao curtir publicação");
    }).always(function() {
        elementoClicado.prop('disabled', false);
    });

}

function atualizarPublicacao() {
    $(this).prop('disabled', true);
    
    const publicacaoId = $(this).data('publicacao-id');
    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "PUT",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
            }
        }).done(function() {
            alert("Publicação atualizada com sucesso");
        }).fail(function(erro) {
            console.log(erro);
            alert("Erro ao atualizar publicação");
        }).always(function() {
            $('#atualizar-publicacao').prop('disabled', false);
        });
    }