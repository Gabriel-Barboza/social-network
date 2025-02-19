$('#nova-publicacao').on('submit', criarpublicacao); 
$(document).on('click', '.curtir-publicacao', curtirpublicacao);
$(document).on('click', '.descurtir-publicacao', descurtirpublicacao);
$('#atualizar-publicacao').on('click', atualizarPublicacao);
$('.deletar-publicacao').on('click', deletarPublicacao);
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
        window.location = "/home";
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao criar a publicação!", "error");
    })
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

    }).fail(function() {
        Swal.fire("Ops...", "Erro ao curtir a publicação!", "error");
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

    }).fail(function() {
        Swal.fire("Ops...", "Erro ao descurtir a publicação!", "error");
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
            Swal.fire('Sucesso!', 'Publicação criada com sucesso!', 'success')
                .then(function() {
                    window.location = "/home";
                })
        }).fail(function() {
            Swal.fire("Ops...", "Erro ao editar a publicação!", "error");
        }).always(function() {
            $('#atualizar-publicacao').prop('disabled', false);
        })
    }

    function deletarPublicacao(evento) {
        evento.preventDefault();
        
        Swal.fire({
            title: "Atenção!",
            text: "Tem certeza que deseja excluir essa publicação? Essa ação é irreversível!",
            showCancelButton: true,
            cancelButtonText: "Cancelar",
            icon: "warning"
        }).then(function(confirmacao) {
            if (!confirmacao.value) return;
    
            const elementoClicado = $(evento.target);
            const publicacao = elementoClicado.closest('div')
            const publicacaoId = publicacao.data('publicacao-id');
        
            elementoClicado.prop('disabled', true);
        
            $.ajax({
                url: `/publicacoes/${publicacaoId}`,
                method: "DELETE"
            }).done(function() {
                publicacao.fadeOut("slow", function() {
                    $(this).remove();
                });
            }).fail(function() {
                Swal.fire("Ops...", "Erro ao excluir a publicação!", "error");
        });});
    }