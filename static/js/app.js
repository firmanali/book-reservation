function Prompt(){
    let toast = function (c) {
        const {
            msg = "",
            icon = "success",
            position = "top-end",
        } = c;
        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            position: position,
            icon: icon,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({
            // icon: c.icon,
            // title: c.msg,
        })
    }

    let success = function (c) {
        const {
            title = "",
            text = "",
            footer = "",
        } = c;
        Swal.fire({
            icon: 'success',
            title: title,
            text: text,
            footer: footer,
        })
    }

    let error = function(c) {
        const {
            title = "",
            text = "",
            footer = "",
        } = c;
        Swal.fire({
            icon: 'error',
            title: title,
            text: text,
            footer: footer,
        })
    }

    async function custom(c) {
        const {
            msg = "",
            title = "",
            icon = "",
            showConfirmButton = true,
            showCancelButton = true,
        } = c;

        const { value: result } = await Swal.fire({
            title: title,
            html: msg,
            icon: icon,
            showConfirmButton: showConfirmButton,
            backdrop: false,
            focusConfirm: false,
            showCancelButton: showCancelButton,
            willOpen: () => {
                if (c.willOpen !== undefined) {
                    c.willOpen();
                }
            },
            preConfirm: () => {
                return [
                    document.getElementById('start').value,
                    document.getElementById('end').value
                ]
            },
            didOpen: () => {
                if(c.didOpen !== undefined){
                    c.didOpen();
                }
            }
        })

        if (result) {
            if(result.dismiss!== Swal.DismissReason.cancel) {
                if(result.value!== ""){
                    if(c.callback!== undefined){
                        c.callback(result);
                    }
                } else {
                    c.callback(false);
                }
            } else {
                c.callback(false);
            }
        }
    }

    return {
        toast :toast,
        success :success,
        error: error,
        custom: custom,
    }
}