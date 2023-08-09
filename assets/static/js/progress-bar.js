(function($) {
  'use strict';
  // ProgressBar JS Starts Here

  if ($('#circleProgress1').length) {
    var bar = new ProgressBar.Circle(circleProgress1, {
      color: '#000',
      // This has to be the same size as the maximum width to
      // prevent clipping
      strokeWidth: 8,
      trailWidth: 8,
      easing: 'easeInOut',
      duration: 1400,
      text: { 
        autoStyleContainer: false
      },
      from: {
        color: '#aaa',
        width: 8,
        radius: 100,
      },
      to: {
        color: '#677ae4',
        width: 8
      },
      // Set default step function for all animate calls
      step: function(state, circle) {
        circle.path.setAttribute('stroke', state.color);
        circle.path.setAttribute('stroke-width', state.width);

        var value = Math.round(circle.value() * 100);
        if (value === 0) {
          circle.setText('');
        } else {
          circle.setText(value);
        }

      }
    });

    bar.text.style.fontSize = '1rem';
    bar.animate(.34); // Number from 0.0 to 1.0
  }

  if ($('#circleProgress1dark').length) {
    var bar = new ProgressBar.Circle(circleProgress1dark, {
      color: '#fff',
      // This has to be the same size as the maximum width to
      // prevent clipping
      strokeWidth: 8,
      trailWidth: 8,
      easing: 'easeInOut',
      duration: 1400,
      text: { 
        autoStyleContainer: false
      },
      from: {
        color: '#3A61F6',
        width: 8,
        radius: 100,
      },
      to: {
        color: '#aaa',
        width: 8
      },
      // Set default step function for all animate calls
      step: function(state, circle) {
        circle.path.setAttribute('stroke', state.color);
        circle.path.setAttribute('stroke-width', state.width);

        var value = Math.round(circle.value() * 100);
        if (value === 0) {
          circle.setText('');
        } else {
          circle.setText(value);
        }

      }
    });

    bar.text.style.fontSize = '1rem';
    bar.animate(.34); // Number from 0.0 to 1.0
  }

  if ($('#circleProgress2').length) {
    var bar = new ProgressBar.Circle(circleProgress2, {
      color: '#000',
      // This has to be the same size as the maximum width to
      // prevent clipping
      strokeWidth: 8,
      trailWidth: 8,
      easing: 'easeInOut',
      duration: 1400,
      text: {
        autoStyleContainer: false
      },
      from: {
        color: '#aaa',
        width: 8
      },
      to: {
        color: '#46c35f',
        width: 8
      },
      // Set default step function for all animate calls
      step: function(state, circle) {
        circle.path.setAttribute('stroke', state.color);
        circle.path.setAttribute('stroke-width', state.width);

        var value = Math.round(circle.value() * 100);
        if (value === 0) {
          circle.setText('');
        } else {
          circle.setText(value);
        }

      }
    });

    bar.text.style.fontSize = '1rem';
    bar.animate(.54); // Number from 0.0 to 1.0
  }
  if ($('#circleProgress2dark').length) {
    var bar = new ProgressBar.Circle(circleProgress2dark, {
      color: '#fff',
      // This has to be the same size as the maximum width to
      // prevent clipping
      strokeWidth: 8,
      trailWidth: 8,
      easing: 'easeInOut',
      duration: 1400,
      text: {
        autoStyleContainer: false
      },
      from: {
        color: '#34B1AA',
        width: 8
      },
      to: {
        color: '#34B1AA',
        width: 8
      },
      // Set default step function for all animate calls
      step: function(state, circle) {
        circle.path.setAttribute('stroke', state.color);
        circle.path.setAttribute('stroke-width', state.width);

        var value = Math.round(circle.value() * 100);
        if (value === 0) {
          circle.setText('');
        } else {
          circle.setText(value);
        }

      }
    });

    bar.text.style.fontSize = '1rem';
    bar.animate(.54); // Number from 0.0 to 1.0
  }

  if ($('#circleProgress3').length) {
    var bar = new ProgressBar.Circle(circleProgress3, {
      color: '#000',
      // This has to be the same size as the maximum width to
      // prevent clipping
      strokeWidth: 8,
      trailWidth: 8,
      easing: 'easeInOut',
      duration: 1400,
      text: {
        autoStyleContainer: false
      },
      from: {
        color: '#aaa',
        width: 8
      },
      to: {
        color: '#f96868',
        width: 8
      },
      // Set default step function for all animate calls
      step: function(state, circle) {
        circle.path.setAttribute('stroke', state.color);
        circle.path.setAttribute('stroke-width', state.width);

        var value = Math.round(circle.value() * 100);
        if (value === 0) {
          circle.setText('');
        } else {
          circle.setText(value);
        }

      }
    });

    bar.text.style.fontSize = '1rem';
    bar.animate(.45); // Number from 0.0 to 1.0
  }
  if ($('#circleProgress3dark').length) {
    var bar = new ProgressBar.Circle(circleProgress3dark, {
      color: '#fff',
      // This has to be the same size as the maximum width to
      // prevent clipping
      strokeWidth: 8,
      trailWidth: 8,
      easing: 'easeInOut',
      duration: 1400,
      text: {
        autoStyleContainer: false
      },
      from: {
        color: '#896C00',
        width: 8
      },
      to: {
        color: '#f96868',
        width: 8
      },
      // Set default step function for all animate calls
      step: function(state, circle) {
        circle.path.setAttribute('stroke', state.color);
        circle.path.setAttribute('stroke-width', state.width);

        var value = Math.round(circle.value() * 100);
        if (value === 0) {
          circle.setText('');
        } else {
          circle.setText(value);
        }

      }
    });

    bar.text.style.fontSize = '1rem';
    bar.animate(.45); // Number from 0.0 to 1.0
  }

  if ($('#circleProgress4').length) {
    var bar = new ProgressBar.Circle(circleProgress4, {
      color: '#000',
      // This has to be the same size as the maximum width to
      // prevent clipping
      strokeWidth: 8,
      trailWidth: 8,
      easing: 'easeInOut',
      duration: 1400,
      text: {
        autoStyleContainer: false
      },
      from: {
        color: '#aaa',
        width: 8
      },
      to: {
        color: '#f2a654',
        width: 8
      },
      // Set default step function for all animate calls
      step: function(state, circle) {
        circle.path.setAttribute('stroke', state.color);
        circle.path.setAttribute('stroke-width', state.width);

        var value = Math.round(circle.value() * 100);
        if (value === 0) {
          circle.setText('');
        } else {
          circle.setText(value);
        }

      }
    });

    bar.text.style.fontSize = '1rem';
    bar.animate(.27); // Number from 0.0 to 1.0
  }
  if ($('#circleProgress4dark').length) {
    var bar = new ProgressBar.Circle(circleProgress4dark, {
      color: '#fff',
      // This has to be the same size as the maximum width to
      // prevent clipping
      strokeWidth: 8,
      trailWidth: 8,
      easing: 'easeInOut',
      duration: 1400,
      text: {
        autoStyleContainer: false
      },
      from: {
        color: '#DEB103',
        width: 8
      },
      to: {
        color: '#DEB103',
        width: 8
      },
      // Set default step function for all animate calls
      step: function(state, circle) {
        circle.path.setAttribute('stroke', state.color);
        circle.path.setAttribute('stroke-width', state.width);

        var value = Math.round(circle.value() * 100);
        if (value === 0) {
          circle.setText('');
        } else {
          circle.setText(value);
        }

      }
    });

    bar.text.style.fontSize = '1rem';
    bar.animate(.27); // Number from 0.0 to 1.0
  }

  if ($('#circleProgress5').length) {
    var bar = new ProgressBar.Circle(circleProgress5, {
      color: '#000',
      // This has to be the same size as the maximum width to
      // prevent clipping
      strokeWidth: 8,
      trailWidth: 8,
      easing: 'easeInOut',
      duration: 1400,
      text: {
        autoStyleContainer: false
      },
      from: {
        color: '#aaa',
        width: 8
      },
      to: {
        color: '#57c7d4',
        width: 8
      },
      // Set default step function for all animate calls
      step: function(state, circle) {
        circle.path.setAttribute('stroke', state.color);
        circle.path.setAttribute('stroke-width', state.width);

        var value = Math.round(circle.value() * 100);
        if (value === 0) {
          circle.setText('');
        } else {
          circle.setText(value);
        }

      }
    });

    bar.text.style.fontSize = '1rem';
    bar.animate(.67); // Number from 0.0 to 1.0
  }
  if ($('#circleProgress5dark').length) {
    var bar = new ProgressBar.Circle(circleProgress5dark, {
      color: '#fff',
      // This has to be the same size as the maximum width to
      // prevent clipping
      strokeWidth: 8,
      trailWidth: 8,
      easing: 'easeInOut',
      duration: 1400,
      text: {
        autoStyleContainer: false
      },
      from: {
        color: '#F95F53',
        width: 8
      },
      to: {
        color: '#F95F53',
        width: 8
      },
      // Set default step function for all animate calls
      step: function(state, circle) {
        circle.path.setAttribute('stroke', state.color);
        circle.path.setAttribute('stroke-width', state.width);

        var value = Math.round(circle.value() * 100);
        if (value === 0) {
          circle.setText('');
        } else {
          circle.setText(value);
        }

      }
    });

    bar.text.style.fontSize = '1rem';
    bar.animate(.67); // Number from 0.0 to 1.0
  }

  if ($('#circleProgress6').length) {
    var bar = new ProgressBar.Circle(circleProgress6, {
      color: '#000',
      // This has to be the same size as the maximum width to
      // prevent clipping
      strokeWidth: 8,
      trailWidth: 8,
      easing: 'easeInOut',
      duration: 1400,
      text: {
        autoStyleContainer: false
      },
      from: {
        color: '#aaa',
        width: 8
      },
      to: {
        color: '#2a2e3b',
        width: 8
      },
      // Set default step function for all animate calls
      step: function(state, circle) {
        circle.path.setAttribute('stroke', state.color);
        circle.path.setAttribute('stroke-width', state.width);

        var value = Math.round(circle.value() * 100);
        if (value === 0) {
          circle.setText('');
        } else {
          circle.setText(value);
        }

      }
    });

    bar.text.style.fontSize = '1rem';
    bar.animate(.95); // Number from 0.0 to 1.0
  }
  if ($('#circleProgress6dark').length) {
    var bar = new ProgressBar.Circle(circleProgress6dark, {
      color: '#fff',
      // This has to be the same size as the maximum width to
      // prevent clipping
      strokeWidth: 8,
      trailWidth: 8,
      easing: 'easeInOut',
      duration: 1400,
      text: {
        autoStyleContainer: false
      },
      from: {
        color: '#a3a4a5',
        width: 8
      },
      to: {
        color: '#a3a4a5',
        width: 8
      },
      // Set default step function for all animate calls
      step: function(state, circle) {
        circle.path.setAttribute('stroke', state.color);
        circle.path.setAttribute('stroke-width', state.width);

        var value = Math.round(circle.value() * 100);
        if (value === 0) {
          circle.setText('');
        } else {
          circle.setText(value);
        }

      }
    });

    bar.text.style.fontSize = '1rem';
    bar.animate(.95); // Number from 0.0 to 1.0
  }

})(jQuery);