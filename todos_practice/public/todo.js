(function ($) {
  'use strict';
  $(function () {
    var todoListItem = $('.todo-list');
    var todoListInput = $('.todo-list-input');
    $('.todo-list-add-btn').on("click", function (event) {
      event.preventDefault();

      var item = $(this).prevAll('.todo-list-input').val();

      if (item) {
        $.post("/todos", { name: item }, addItem)
        todoListInput.val("");
      }

    });

    $.get("/todos", function (items) {
      items.forEach(element => {
        addItem(element)
      })
    })

    var addItem = function (item) {
      if (item.completed) {
        todoListItem.append("<li class='completed' id='" + item.id + "'><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' />" + item.name + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
      } else {
        todoListItem.append("<li id='" + item.id + "'><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' />" + item.name + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
      }
    }

    todoListItem.on('change', '.checkbox', function () {

      var $self = $(this)
      var id = $self.closest("li").attr("id")

      var complete = true

      if ($self.attr('checked')) {
        complete = false
      }

      $.get("/complete-todo/" + id + "?complete=" + complete, function (rst) {
        if (complete) {
          $self.attr('checked', 'checked');
        } else {
          $self.removeAttr('checked');
        }
      })

      $(this).closest("li").toggleClass('completed');

    });

    todoListItem.on('click', '.remove', function () {
      var $self = $(this)
      var id = $self.closest("li").attr("id")
      $.ajax({
        url: "/todos/" + id,
        type: "DELETE",
        success: function (data) {
          if (data.success) {
            $self.parent().remove();
          }
        }
      })

    });

  });
})(jQuery);
