(function ($) {
  'use strict';
  $(function () {
    var todoListItem = $('.todo-list');
    var todoListInput = $('.todo-list-input');
    $('.todo-list-add-btn').on("click", function (event) {
      event.preventDefault();

      var item = $(this).prevAll('.todo-list-input').val();

      if (item) {
        $.post("/todos", { name: item }, addItem);
        todoListInput.val("");
      }

    });

    var addItem = function (item) {
      if (item.completed) {
        todoListItem.append("<li class='completed' id='" + item.id + "'><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' checked='checked' />" + item.name + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
      } else {
        todoListItem.append("<li id='" + item.id + "'><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' />" + item.name + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
      }
    }

    $.get("/todos", function (items) {
      items.forEach(element => {
        addItem(element)
      })
    });

    todoListItem.on('change', '.checkbox', function () {

      var id = $(this).closest("li").attr("id")
      // 값을 실행될때 this가 바뀌기 때문에 저장해놓고 추후에 사용
      var $self = $(this)
      var complete = true;
      if ($(this).attr("checked")) {
        complete = false;
      }
      $.get("complete-todo/" + id + "?complete=" + complete, function (data) {
        if (complete) {
          $self.attr('checked', 'checked');
        } else {
          $self.removeAttr('checked');
        }

        $self.closest("li").toggleClass('completed');
      })

      // if ($(this).attr('checked')) {
      //   $(this).removeAttr('checked');
      // } else {
      //   $(this).attr('checked', 'checked');
      // }

      // $(this).closest("li").toggleClass('completed');


    });

    todoListItem.on('click', '.remove', function () {
      // $(this).parent().remove();
      // 서버에 요청
      // url: todos/{id}, method: DELETE
      // this : remove button
      var id = $(this).closest("li").attr("id")
      var $self = $(this)
      $.ajax({
        // request object 만들기
        url: "todos/" + id,
        type: "DELETE",
        success: function (data) {
          // 성공했을 때 수행할 것
          if (data.success) {
            $self.parent().remove();
          }
        }
      })
    });

  });
})(jQuery);
