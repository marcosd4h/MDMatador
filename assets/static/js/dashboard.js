(function($) {
  'use strict';
  $(function() {
    if ($("#performaneLine").length) {
      var graphGradient = document.getElementById("performaneLine").getContext('2d');
      var graphGradient2 = document.getElementById("performaneLine").getContext('2d');
      var saleGradientBg = graphGradient.createLinearGradient(5, 0, 5, 100);
      saleGradientBg.addColorStop(0, 'rgba(26, 115, 232, 0.18)');
      saleGradientBg.addColorStop(1, 'rgba(26, 115, 232, 0.02)');
      var saleGradientBg2 = graphGradient2.createLinearGradient(100, 0, 50, 150);
      saleGradientBg2.addColorStop(0, 'rgba(0, 208, 255, 0.19)');
      saleGradientBg2.addColorStop(1, 'rgba(0, 208, 255, 0.03)');
      var salesTopData = {
          labels: ["SUN","sun", "MON", "mon", "TUE","tue", "WED", "wed", "THU", "thu", "FRI", "fri", "SAT"],
          datasets: [{
              label: 'This week',
              data: [50, 110, 60, 290, 200, 115, 130, 170, 90, 210, 240, 280, 200],
              backgroundColor: saleGradientBg,
              borderColor: [
                  '#1F3BB3',
              ],
              borderWidth: 1.5,
              fill: true, // 3: no fill
              pointBorderWidth: 1,
              pointRadius: [4, 4, 4, 4, 4,4, 4, 4, 4, 4,4, 4, 4],
              pointHoverRadius: [2, 2, 2, 2, 2,2, 2, 2, 2, 2,2, 2, 2],
              pointBackgroundColor: ['#1F3BB3)', '#1F3BB3', '#1F3BB3', '#1F3BB3','#1F3BB3)', '#1F3BB3', '#1F3BB3', '#1F3BB3','#1F3BB3)', '#1F3BB3', '#1F3BB3', '#1F3BB3','#1F3BB3)'],
              pointBorderColor: ['#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff',],
          },{
            label: 'Last week',
            data: [30, 150, 190, 250, 120, 150, 130, 20, 30, 15, 40, 95, 180],
            backgroundColor: saleGradientBg2,
            borderColor: [
                '#52CDFF',
            ],
            borderWidth: 1.5,
            fill: true, // 3: no fill
            pointBorderWidth: 1,
            pointRadius: [0, 0, 0, 4, 0],
            pointHoverRadius: [0, 0, 0, 2, 0],
            pointBackgroundColor: ['#52CDFF)', '#52CDFF', '#52CDFF', '#52CDFF','#52CDFF)', '#52CDFF', '#52CDFF', '#52CDFF','#52CDFF)', '#52CDFF', '#52CDFF', '#52CDFF','#52CDFF)'],
              pointBorderColor: ['#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff',],
        }]
      };
  
      var salesTopOptions = {
        responsive: true,
        maintainAspectRatio: false,
          scales: {
              yAxes: [{
                  gridLines: {
                      display: true,
                      drawBorder: false,
                      color:"#F0F0F0",
                      zeroLineColor: '#F0F0F0',
                  },
                  ticks: {
                    beginAtZero: false,
                    autoSkip: true,
                    maxTicksLimit: 4,
                    fontSize: 10,
                    color:"#6B778C"
                  }
              }],
              xAxes: [{
                gridLines: {
                    display: false,
                    drawBorder: false,
                },
                ticks: {
                  beginAtZero: false,
                  autoSkip: true,
                  maxTicksLimit: 7,
                  fontSize: 10,
                  color:"#6B778C"
                }
            }],
          },
          legend:false,
          legendCallback: function (chart) {
            var text = [];
            text.push('<div class="chartjs-legend"><ul>');
            for (var i = 0; i < chart.data.datasets.length; i++) {
              console.log(chart.data.datasets[i]); // see what's inside the obj.
              text.push('<li>');
              text.push('<span style="background-color:' + chart.data.datasets[i].borderColor + '">' + '</span>');
              text.push(chart.data.datasets[i].label);
              text.push('</li>');
            }
            text.push('</ul></div>');
            return text.join("");
          },
          
          elements: {
              line: {
                  tension: 0.4,
              }
          },
          tooltips: {
              backgroundColor: 'rgba(31, 59, 179, 1)',
          }
      }
      var salesTop = new Chart(graphGradient, {
          type: 'line',
          data: salesTopData,
          options: salesTopOptions
      });
      document.getElementById('performance-line-legend').innerHTML = salesTop.generateLegend();
    }
    if ($("#performaneLine-dark").length) {
      var graphGradient = document.getElementById("performaneLine-dark").getContext('2d');
      var graphGradient2 = document.getElementById("performaneLine-dark").getContext('2d');
      var saleGradientBg = graphGradient.createLinearGradient(5, 0, 5, 100);
      saleGradientBg.addColorStop(0, 'rgba(251, 150, 0, 0.18)');
      saleGradientBg.addColorStop(1, 'rgba(251, 150, 0, 0.02)');
      var saleGradientBg2 = graphGradient2.createLinearGradient(10, 0, 0, 150);
      saleGradientBg2.addColorStop(0, 'rgba(251, 255, 255, 0)');
      saleGradientBg2.addColorStop(1, 'rgba(251, 255, 255, 0)');
      var salesTopDataDark = {
          labels: ["SUN","sun", "MON", "mon", "TUE","tue", "WED", "wed", "THU", "thu", "FRI", "fri", "SAT"],
          datasets: [{
              label: '# of Votes',
              data: [50, 110, 60, 290, 200, 115, 130, 170, 90, 210, 240, 280, 200],
              backgroundColor: saleGradientBg,
              borderColor: [
                  '#F29F67',
              ],
              borderWidth: 1.5,
              fill: true, // 3: no fill
              pointBorderWidth: 1,
              pointRadius: [4, 4, 4, 4, 4,4, 4, 4, 4, 4,4, 4, 4],
              pointHoverRadius: [2, 2, 2, 2, 2,2, 2, 2, 2, 2,2, 2, 2],
              pointBackgroundColor: ['#F29F67)', '#F29F67', '#F29F67', '#F29F67','#F29F67)', '#F29F67', '#F29F67', '#F29F67','#F29F67)', '#F29F67', '#F29F67', '#F29F67','#F29F67)'],
              pointBorderColor: ['#252730','#252730','#252730','#252730','#252730','#252730','#252730','#252730','#252730','#252730','#252730','#252730','#252730',],
          },{
            label: '# of Votes',
            data: [30, 150, 190, 250, 120, 150, 130, 20, 30, 15, 40, 95, 180],
            backgroundColor: saleGradientBg2,
            borderColor: [
                '#808191',
            ],
            borderWidth: 1.5,
            fill: true, // 3: no fill
            pointBorderWidth: 1,
            pointRadius: [0, 0, 0, 4, 0],
            pointHoverRadius: [0, 0, 0, 2, 0],
            pointBackgroundColor: ['#808191)', '#808191', '#808191', '#808191','#808191)', '#808191', '#808191', '#808191','#808191)', '#808191', '#808191', '#808191','#808191)'],
            pointBorderColor: ['#252730','#252730','#252730','#252730','#252730','#252730','#252730','#252730','#252730','#252730','#252730','#252730','#252730',],
        }]
      };
  
      var salesTopOptionsDark = {
        responsive: true,
        maintainAspectRatio: false,
          scales: {
              yAxes: [{
                  gridLines: {
                      display: true,
                      drawBorder: false,
                      color:"#383A42",
                      zeroLineColor: "#383A42",
                  },
                  ticks: {
                    beginAtZero: false,
                    autoSkip: true,
                    maxTicksLimit: 4,
                    fontSize: 10,
                    color:"#808191"
                  }
              }],
              xAxes: [{
                gridLines: {
                    display: false,
                    drawBorder: false,
                },
                ticks: {
                  beginAtZero: false,
                  autoSkip: true,
                  maxTicksLimit: 7,
                  fontSize: 10,
                  color:"#6B778C"
                }
            }],
          },
          legend:false,
          legendCallback: function (chart) {
            var text = [];
            text.push('<div class="chartjs-legend"><ul>');
            for (var i = 0; i < chart.data.datasets.length; i++) {
              console.log(chart.data.datasets[i]); // see what's inside the obj.
              text.push('<li>');
              text.push('<span style="background-color:' + chart.data.datasets[i].borderColor + '">' + '</span>');
              text.push(chart.data.datasets[i].label);
              text.push('</li>');
            }
            text.push('</ul></div>');
            return text.join("");
          },
          
          elements: {
              line: {
                  tension: 0.4,
              }
          },
          tooltips: {
              backgroundColor: 'rgba(31, 59, 179, 1)',
          }
      }
      var salesTopDark = new Chart(graphGradient, {
          type: 'line',
          data: salesTopDataDark,
          options: salesTopOptionsDark
      });
      document.getElementById('performance-line-legend-dark').innerHTML = salesTopDark.generateLegend();
    }
    if ($("#performaneLineBrown").length) {
      var graphGradient = document.getElementById("performaneLineBrown").getContext('2d');
      var graphGradient2 = document.getElementById("performaneLineBrown").getContext('2d');
      var saleGradientBg = graphGradient.createLinearGradient(5, 0, 5, 100);
      saleGradientBg.addColorStop(0, 'rgba(251, 150, 0, 0.18)');
      saleGradientBg.addColorStop(1, 'rgba(251, 150, 0, 0.02)');
      var saleGradientBg2 = graphGradient2.createLinearGradient(100, 0, 50, 150);
      saleGradientBg2.addColorStop(0, 'rgba(0, 0, 0, 0.18)');
      saleGradientBg2.addColorStop(1, 'rgba(0, 0, 0, 0.01)');
      var salesTopData = {
          labels: ["SUN","sun", "MON", "mon", "TUE","tue", "WED", "wed", "THU", "thu", "FRI", "fri", "SAT"],
          datasets: [{
              label: 'This week',
              data: [50, 110, 60, 290, 200, 115, 130, 170, 90, 210, 240, 280, 200],
              backgroundColor: saleGradientBg,
              borderColor: [
                  '#F29F67',
              ],
              borderWidth: 1.5,
              fill: true, // 3: no fill
              pointBorderWidth: 1,
              pointRadius: [4, 4, 4, 4, 4,4, 4, 4, 4, 4,4, 4, 4],
              pointHoverRadius: [2, 2, 2, 2, 2,2, 2, 2, 2, 2,2, 2, 2],
              pointBackgroundColor: ['#F29F67)', '#F29F67', '#F29F67', '#F29F67','#F29F67)', '#F29F67', '#F29F67', '#F29F67','#F29F67)', '#F29F67', '#F29F67', '#F29F67','#F29F67)'],
              pointBorderColor: ['#FAF8F3','#FAF8F3','#FAF8F3','#FAF8F3','#FAF8F3','#FAF8F3','#FAF8F3','#FAF8F3','#FAF8F3','#FAF8F3','#FAF8F3','#FAF8F3','#FAF8F3',],
          },{
            label: 'Last week',
            data: [30, 150, 190, 250, 120, 150, 130, 20, 30, 15, 40, 95, 180],
            backgroundColor: saleGradientBg2,
            borderColor: [
                '#000000',
            ],
            borderWidth: 1.5,
            fill: true, // 3: no fill
            pointBorderWidth: 1,
            pointRadius: [0, 0, 0, 4, 0],
            pointHoverRadius: [0, 0, 0, 2, 0],
            pointBackgroundColor: ['#000000)', '#000000', '#000000', '#000000','#000000)', '#000000', '#000000', '#000000','#000000)', '#000000', '#000000', '#000000','#000000)'],
              pointBorderColor: ['#FAF8F3','#FAF8F3','#FAF8F3','#FAF8F3','#FAF8F3','#FAF8F3','#FAF8F3','#FAF8F3','#FAF8F3','#FAF8F3','#FAF8F3','#FAF8F3','#FAF8F3',],
        }]
      };
  
      var salesTopOptions = {
        responsive: true,
        maintainAspectRatio: false,
          scales: {
              yAxes: [{
                  gridLines: {
                      display: true,
                      drawBorder: false,
                      color:"#F0F0F0",
                      zeroLineColor: '#F0F0F0',
                  },
                  ticks: {
                    beginAtZero: false,
                    autoSkip: true,
                    maxTicksLimit: 4,
                    fontSize: 10,
                    color:"#6B778C"
                  }
              }],
              xAxes: [{
                gridLines: {
                    display: false,
                    drawBorder: false,
                },
                ticks: {
                  beginAtZero: false,
                  autoSkip: true,
                  maxTicksLimit: 7,
                  fontSize: 10,
                  color:"#6B778C"
                }
            }],
          },
          legend:false,
          legendCallback: function (chart) {
            var text = [];
            text.push('<div class="chartjs-legend"><ul>');
            for (var i = 0; i < chart.data.datasets.length; i++) {
              console.log(chart.data.datasets[i]); // see what's inside the obj.
              text.push('<li>');
              text.push('<span style="background-color:' + chart.data.datasets[i].borderColor + '">' + '</span>');
              text.push(chart.data.datasets[i].label);
              text.push('</li>');
            }
            text.push('</ul></div>');
            return text.join("");
          },
          
          elements: {
              line: {
                  tension: 0.4,
              }
          },
          tooltips: {
              backgroundColor: 'rgba(31, 59, 179, 1)',
          }
      }
      var salesTop = new Chart(graphGradient, {
          type: 'line',
          data: salesTopData,
          options: salesTopOptions
      });
      document.getElementById('performance-line-legend').innerHTML = salesTop.generateLegend();
    }
    if ($("#performaneLinePurple").length) {
      var graphGradient = document.getElementById("performaneLinePurple").getContext('2d');
      var graphGradient2 = document.getElementById("performaneLinePurple").getContext('2d');
      var saleGradientBg = graphGradient.createLinearGradient(5, 0, 5, 100);
      saleGradientBg.addColorStop(0, 'rgba(0, 123, 255, 0.18)');
      saleGradientBg.addColorStop(1, 'rgba(0, 123, 255, 0.02)');
      var saleGradientBg2 = graphGradient2.createLinearGradient(100, 0, 50, 150);
      saleGradientBg2.addColorStop(0, 'rgba(86, 11, 208, 0.12)');
      saleGradientBg2.addColorStop(1, 'rgba(86, 11, 208, 0.03)');
      var salesTopData = {
          labels: ["SUN","sun", "MON", "mon", "TUE","tue", "WED", "wed", "THU", "thu", "FRI", "fri", "SAT"],
          datasets: [{
              label: 'This week',
              data: [50, 110, 60, 290, 200, 115, 130, 170, 90, 210, 240, 280, 200],
              backgroundColor: saleGradientBg,
              borderColor: [
                  '#006CFF',
              ],
              borderWidth: 1.5,
              fill: true, // 3: no fill
              pointBorderWidth: 1,
              pointRadius: [4, 4, 4, 4, 4,4, 4, 4, 4, 4,4, 4, 4],
              pointHoverRadius: [2, 2, 2, 2, 2,2, 2, 2, 2, 2,2, 2, 2],
              pointBackgroundColor: ['#1F3BB3)', '#1F3BB3', '#1F3BB3', '#1F3BB3','#1F3BB3)', '#1F3BB3', '#1F3BB3', '#1F3BB3','#1F3BB3)', '#1F3BB3', '#1F3BB3', '#1F3BB3','#1F3BB3)'],
              pointBorderColor: ['#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF',],
          },{
            label: 'Last week',
            data: [30, 150, 190, 250, 120, 150, 130, 20, 30, 15, 40, 95, 180],
            backgroundColor: saleGradientBg2,
            borderColor: [
                '#7B20C7',
            ],
            borderWidth: 1.5,
            fill: true, // 3: no fill
            pointBorderWidth: 1,
            pointRadius: [0, 0, 0, 4, 0],
            pointHoverRadius: [0, 0, 0, 2, 0],
            pointBackgroundColor: ['#52CDFF)', '#52CDFF', '#52CDFF', '#52CDFF','#52CDFF)', '#52CDFF', '#52CDFF', '#52CDFF','#52CDFF)', '#52CDFF', '#52CDFF', '#52CDFF','#52CDFF)'],
              pointBorderColor: ['#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF',],
        }]
      };
  
      var salesTopOptions = {
        responsive: true,
        maintainAspectRatio: false,
          scales: {
              yAxes: [{
                  gridLines: {
                      display: true,
                      drawBorder: false,
                      color:"#F0F0F0",
                      zeroLineColor: '#F0F0F0',
                  },
                  ticks: {
                    beginAtZero: false,
                    autoSkip: true,
                    maxTicksLimit: 4,
                    fontSize: 10,
                    color:"#6B778C"
                  }
              }],
              xAxes: [{
                gridLines: {
                    display: false,
                    drawBorder: false,
                },
                ticks: {
                  beginAtZero: false,
                  autoSkip: true,
                  maxTicksLimit: 7,
                  fontSize: 10,
                  color:"#6B778C"
                }
            }],
          },
          legend:false,
          legendCallback: function (chart) {
            var text = [];
            text.push('<div class="chartjs-legend"><ul>');
            for (var i = 0; i < chart.data.datasets.length; i++) {
              console.log(chart.data.datasets[i]); // see what's inside the obj.
              text.push('<li>');
              text.push('<span style="background-color:' + chart.data.datasets[i].borderColor + '">' + '</span>');
              text.push(chart.data.datasets[i].label);
              text.push('</li>');
            }
            text.push('</ul></div>');
            return text.join("");
          },
          
          elements: {
              line: {
                  tension: 0.4,
              }
          },
          tooltips: {
              backgroundColor: 'rgba(31, 59, 179, 1)',
          }
      }
      var salesTop = new Chart(graphGradient, {
          type: 'line',
          data: salesTopData,
          options: salesTopOptions
      });
      document.getElementById('performance-line-legend').innerHTML = salesTop.generateLegend();
    }
    if ($("#performaneLinePurple-dark").length) {
      var graphGradient = document.getElementById("performaneLinePurple-dark").getContext('2d');
      var graphGradient2 = document.getElementById("performaneLinePurple-dark").getContext('2d');
      var saleGradientBg = graphGradient.createLinearGradient(5, 0, 5, 100);
      saleGradientBg.addColorStop(0, 'rgba(174, 77, 255, 0.19)');
      saleGradientBg.addColorStop(1, 'rgba(157, 53, 244, 0.03)');
      var saleGradientBg2 = graphGradient2.createLinearGradient(100, 0, 50, 150);
      saleGradientBg2.addColorStop(0, 'rgba(0, 163, 255, 0.1)');
      saleGradientBg2.addColorStop(1, 'rgba(0, 163, 255, 0.02)');
      var salesTopData = {
          labels: ["SUN","sun", "MON", "mon", "TUE","tue", "WED", "wed", "THU", "thu", "FRI", "fri", "SAT"],
          datasets: [{
              label: 'This week',
              data: [50, 110, 60, 290, 200, 115, 130, 170, 90, 210, 240, 280, 200],
              backgroundColor: saleGradientBg,
              borderColor: [
                  '#BE70FF',
              ],
              borderWidth: 1.5,
              fill: true, // 3: no fill
              pointBorderWidth: 1,
              pointRadius: [4, 4, 4, 4, 4,4, 4, 4, 4, 4,4, 4, 4],
              pointHoverRadius: [2, 2, 2, 2, 2,2, 2, 2, 2, 2,2, 2, 2],
              pointBackgroundColor: ['#1F3BB3)', '#1F3BB3', '#1F3BB3', '#1F3BB3','#1F3BB3)', '#1F3BB3', '#1F3BB3', '#1F3BB3','#1F3BB3)', '#1F3BB3', '#1F3BB3', '#1F3BB3','#1F3BB3)'],
              pointBorderColor: ['#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF',],
          },{
            label: 'Last week',
            data: [30, 150, 190, 250, 120, 150, 130, 20, 30, 15, 40, 95, 180],
            backgroundColor: saleGradientBg2,
            borderColor: [
                '#00A3FF',
            ],
            borderWidth: 1.5,
            fill: true, // 3: no fill
            pointBorderWidth: 1,
            pointRadius: [0, 0, 0, 4, 0],
            pointHoverRadius: [0, 0, 0, 2, 0],
            pointBackgroundColor: ['#52CDFF)', '#52CDFF', '#52CDFF', '#52CDFF','#52CDFF)', '#52CDFF', '#52CDFF', '#52CDFF','#52CDFF)', '#52CDFF', '#52CDFF', '#52CDFF','#52CDFF)'],
            pointBorderColor: ['#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF','#FFFFFF',],
        }]
      };
  
      var salesTopOptions = {
        responsive: true,
        maintainAspectRatio: false,
          scales: {
              yAxes: [{
                  gridLines: {
                      display: true,
                      drawBorder: false,
                      color:"rgba(240,240,240, .2)",
                      zeroLineColor: 'rgba(240,240,240, .2)',
                  },
                  ticks: {
                    beginAtZero: false,
                    autoSkip: true,
                    maxTicksLimit: 4,
                    fontSize: 10,
                    color:"rgba(240,240,240, .2)"
                  }
              }],
              xAxes: [{
                gridLines: {
                    display: false,
                    drawBorder: false,
                },
                ticks: {
                  beginAtZero: false,
                  autoSkip: true,
                  maxTicksLimit: 7,
                  fontSize: 10,
                  color:"rgba(240,240,240, .2)"
                }
            }],
          },
          legend:false,
          legendCallback: function (chart) {
            var text = [];
            text.push('<div class="chartjs-legend"><ul>');
            for (var i = 0; i < chart.data.datasets.length; i++) {
              console.log(chart.data.datasets[i]); // see what's inside the obj.
              text.push('<li>');
              text.push('<span style="background-color:' + chart.data.datasets[i].borderColor + '">' + '</span>');
              text.push(chart.data.datasets[i].label);
              text.push('</li>');
            }
            text.push('</ul></div>');
            return text.join("");
          },
          
          elements: {
              line: {
                  tension: 0.4,
              }
          },
          tooltips: {
              backgroundColor: 'rgba(31, 59, 179, 1)',
          }
      }
      var salesTop = new Chart(graphGradient, {
          type: 'line',
          data: salesTopData,
          options: salesTopOptions
      });
      document.getElementById('performance-line-legend').innerHTML = salesTop.generateLegend();
    }
    if ($("#salesTrend").length) {
      var graphGradient = document.getElementById("salesTrend").getContext('2d');
      var graphGradient2 = document.getElementById("salesTrend").getContext('2d');
      var saleGradientBg = graphGradient.createLinearGradient(5, 0, 5, 100);
      saleGradientBg.addColorStop(0, 'rgba(42, 33, 186, 0.2)');
      saleGradientBg.addColorStop(1, 'rgba(42, 33, 186, 0)');
      var saleGradientBg2 = graphGradient2.createLinearGradient(100, 0, 50, 150);
      saleGradientBg2.addColorStop(0, 'rgba(0, 205, 255, 0.2)');
      saleGradientBg2.addColorStop(1, 'rgba(0, 205, 255, 0)');
      var salesTopData = {
          labels: ["SUN","sun", "MON", "mon", "TUE","tue", "WED", "wed", "THU", "thu", "FRI", "fri", "SAT"],
          datasets: [{
              label: 'Online Payment',
              data: [50, 110, 60, 290, 200, 115, 130, 170, 90, 210, 240, 280, 200],
              backgroundColor: saleGradientBg,
              borderColor: [
                  '#2A21BA',
              ],
              borderWidth: 1.5,
              fill: true, // 3: no fill
              pointBorderWidth: 1,
              pointRadius: [4, 4, 4, 4, 4,4, 4, 4, 4, 4,4, 4, 4],
              pointHoverRadius: [2, 2, 2, 2, 2,2, 2, 2, 2, 2,2, 2, 2],
              pointBackgroundColor: ['#1F3BB3)', '#1F3BB3', '#1F3BB3', '#1F3BB3','#1F3BB3)', '#1F3BB3', '#1F3BB3', '#1F3BB3','#1F3BB3)', '#1F3BB3', '#1F3BB3', '#1F3BB3','#1F3BB3)'],
              pointBorderColor: ['#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff',],
          },{
            label: 'Offline Sales',
            data: [30, 150, 190, 250, 120, 150, 130, 20, 30, 15, 40, 95, 180],
            backgroundColor: saleGradientBg2,
            borderColor: [
                '#52CDFF',
            ],
            borderWidth: 1.5,
            fill: true, // 3: no fill
            pointBorderWidth: 1,
            pointRadius: [0, 0, 0, 4, 0],
            pointHoverRadius: [0, 0, 0, 2, 0],
            pointBackgroundColor: ['#00CDFF)', '#00CDFF', '#00CDFF', '#00CDFF','#00CDFF)', '#00CDFF', '#00CDFF', '#00CDFF','#00CDFF)', '#00CDFF', '#00CDFF', '#00CDFF','#00CDFF)'],
              pointBorderColor: ['#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff',],
        }]
      };
  
      var salesTopOptions = {
        responsive: true,
        maintainAspectRatio: false,
          scales: {
              yAxes: [{
                  gridLines: {
                      display: true,
                      drawBorder: false,
                      color:"#F0F0F0",
                      zeroLineColor: '#F0F0F0',
                  },
                  ticks: {
                    beginAtZero: false,
                    autoSkip: true,
                    maxTicksLimit: 4,
                    fontSize: 10,
                    color:"#6B778C"
                  }
              }],
              xAxes: [{
                gridLines: {
                    display: false,
                    drawBorder: false,
                },
                ticks: {
                  beginAtZero: false,
                  autoSkip: true,
                  maxTicksLimit: 7,
                  fontSize: 10,
                  color:"#6B778C"
                }
            }],
          },
          legend:false,
          legendCallback: function (chart) {
            var text = [];
            text.push('<div class="chartjs-legend"><ul>');
            for (var i = 0; i < chart.data.datasets.length; i++) {
              console.log(chart.data.datasets[i]); // see what's inside the obj.
              text.push('<li>');
              text.push('<span style="background-color:' + chart.data.datasets[i].borderColor + '">' + '</span>');
              text.push(chart.data.datasets[i].label);
              text.push('</li>');
            }
            text.push('</ul></div>');
            return text.join("");
          },
          
          elements: {
              line: {
                  tension: 0.4,
              }
          },
          tooltips: {
              backgroundColor: 'rgba(31, 59, 179, 1)',
          }
      }
      var salesTop = new Chart(graphGradient, {
          type: 'line',
          data: salesTopData,
          options: salesTopOptions
      });
      document.getElementById('sales-trend-legend').innerHTML = salesTop.generateLegend();
    }
    if ($("#status-summary").length) {
      var statusSummaryChartCanvas = document.getElementById("status-summary").getContext('2d');;
      var statusData = {
          labels: ["SUN", "MON", "TUE", "WED", "THU", "FRI"],
          datasets: [{
              label: '# of Votes',
              data: [50, 68, 70, 10, 12, 80],
              backgroundColor: "#ffcc00",
              borderColor: [
                  '#01B6A0',
              ],
              borderWidth: 2,
              fill: false, // 3: no fill
              pointBorderWidth: 0,
              pointRadius: [0, 0, 0, 0, 0, 0],
              pointHoverRadius: [0, 0, 0, 0, 0, 0],
          }]
      };
  
      var statusOptions = {
        responsive: true,
        maintainAspectRatio: false,
          scales: {
              yAxes: [{
                display:false,
                  gridLines: {
                      display: false,
                      drawBorder: false,
                      color:"#F0F0F0"
                  },
                  ticks: {
                    beginAtZero: false,
                    autoSkip: true,
                    maxTicksLimit: 4,
                    fontSize: 10,
                    color:"#6B778C"
                  }
              }],
              xAxes: [{
                display:false,
                gridLines: {
                    display: false,
                    drawBorder: false,
                },
                ticks: {
                  beginAtZero: false,
                  autoSkip: true,
                  maxTicksLimit: 7,
                  fontSize: 10,
                  color:"#6B778C"
                }
            }],
          },
          legend:false,
          
          elements: {
              line: {
                  tension: 0.4,
              }
          },
          tooltips: {
              backgroundColor: 'rgba(31, 59, 179, 1)',
          }
      }
      var statusSummaryChart = new Chart(statusSummaryChartCanvas, {
          type: 'line',
          data: statusData,
          options: statusOptions
      });
    }
    if ($("#status-summary-purple").length) {
      var statusSummaryChartCanvas = document.getElementById("status-summary-purple").getContext('2d');;
      var statusData = {
          labels: ["SUN", "MON", "TUE", "WED", "THU", "FRI"],
          datasets: [{
              label: '# of Votes',
              data: [50, 68, 70, 10, 12, 80],
              backgroundColor: "#ffcc00",
              borderColor: [
                  '#00CCCC',
              ],
              borderWidth: 2,
              fill: false, // 3: no fill
              pointBorderWidth: 0,
              pointRadius: [0, 0, 0, 0, 0, 0],
              pointHoverRadius: [0, 0, 0, 0, 0, 0],
          }]
      };
  
      var statusOptions = {
        responsive: true,
        maintainAspectRatio: false,
          scales: {
              yAxes: [{
                display:false,
                  gridLines: {
                      display: false,
                      drawBorder: false,
                      color:"#F0F0F0"
                  },
                  ticks: {
                    beginAtZero: false,
                    autoSkip: true,
                    maxTicksLimit: 4,
                    fontSize: 10,
                    color:"#6B778C"
                  }
              }],
              xAxes: [{
                display:false,
                gridLines: {
                    display: false,
                    drawBorder: false,
                },
                ticks: {
                  beginAtZero: false,
                  autoSkip: true,
                  maxTicksLimit: 7,
                  fontSize: 10,
                  color:"#6B778C"
                }
            }],
          },
          legend:false,
          
          elements: {
              line: {
                  tension: 0.4,
              }
          },
          tooltips: {
              backgroundColor: 'rgba(31, 59, 179, 1)',
          }
      }
      var statusSummaryChart = new Chart(statusSummaryChartCanvas, {
          type: 'line',
          data: statusData,
          options: statusOptions
      });
    }
    if ($('#totalVisitors').length) {
      var bar = new ProgressBar.Circle(totalVisitors, {
        color: '#fff',
        // This has to be the same size as the maximum width to
        // prevent clipping
        strokeWidth: 15,
        trailWidth: 15, 
        easing: 'easeInOut',
        duration: 1400,
        text: {
          autoStyleContainer: false
        },
        from: {
          color: '#52CDFF',
          width: 15
        },
        to: {
          color: '#677ae4',
          width: 15
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
  
      bar.text.style.fontSize = '0rem';
      bar.animate(.64); // Number from 0.0 to 1.0
    }
    if ($('#totalVisitorsDark').length) {
      var bar = new ProgressBar.Circle(totalVisitorsDark, {
        color: '#4A4C55',
        // This has to be the same size as the maximum width to
        // prevent clipping
        strokeWidth: 15,
        trailWidth: 15, 
        trailColor: "#4A4C55", 
        easing: 'easeInOut',
        duration: 1400,
        text: {
          autoStyleContainer: false
        },
        from: {
          color: '#3A61F6',
          width: 15,
        },
        to: {
          color: '#3A61F6',
          width: 15
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
  
      bar.text.style.fontSize = '0rem';
      bar.animate(.64); // Number from 0.0 to 1.0
    }
    if ($('#totalVisitorsPurple').length) {
      var bar = new ProgressBar.Circle(totalVisitorsPurple, {
        color: '#fff',
        // This has to be the same size as the maximum width to
        // prevent clipping
        strokeWidth: 15,
        trailWidth: 15, 
        easing: 'easeInOut',
        duration: 1400,
        text: {
          autoStyleContainer: false
        },
        from: {
          color: '#52CDFF',
          width: 15
        },
        to: {
          color: '#006CFF',
          width: 15
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
  
      bar.text.style.fontSize = '0rem';
      bar.animate(.64); // Number from 0.0 to 1.0
    }
    if ($('#visitperday').length) {
      var bar = new ProgressBar.Circle(visitperday, {
        color: '#fff',
        // This has to be the same size as the maximum width to
        // prevent clipping
        strokeWidth: 15,
        trailWidth: 15,
        easing: 'easeInOut',
        duration: 1400,
        text: {
          autoStyleContainer: false
        },
        from: {
          color: '#34B1AA',
          width: 15
        },
        to: {
          color: '#677ae4',
          width: 15
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
  
      bar.text.style.fontSize = '0rem';
      bar.animate(.34); // Number from 0.0 to 1.0
    }
    if ($('#visitperdayPurple').length) {
      var bar = new ProgressBar.Circle(visitperdayPurple, {
        color: '#fff',
        // This has to be the same size as the maximum width to
        // prevent clipping
        strokeWidth: 15,
        trailWidth: 15,
        easing: 'easeInOut',
        duration: 1400,
        text: {
          autoStyleContainer: false
        },
        from: {
          color: '#7B20C7',
          width: 15
        },
        to: {
          color: '#7B20C7',
          width: 15
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
  
      bar.text.style.fontSize = '0rem';
      bar.animate(.34); // Number from 0.0 to 1.0
    }
    if ($('#visitperdayDark').length) {
      var bar = new ProgressBar.Circle(visitperdayDark, {
        color: '#4A4C55',
        // This has to be the same size as the maximum width to
        // prevent clipping
        strokeWidth: 15,
        trailWidth: 15, 
        trailColor: "#4A4C55", 
        easing: 'easeInOut',
        duration: 1400,
        text: {
          autoStyleContainer: false
        },
        from: {
          color: '#04A390',
          width: 15,
        },
        to: {
          color: '#04A390',
          width: 15
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
  
      bar.text.style.fontSize = '0rem';
      bar.animate(.64); // Number from 0.0 to 1.0
    }
    if ($("#marketingOverview").length) {
      var marketingOverviewChart = document.getElementById("marketingOverview").getContext('2d');
      var marketingOverviewData = {
          labels: ["JAN","FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"],
          datasets: [{
            label: 'Last week',
            data: [110, 220, 200, 190, 220, 110, 210, 110, 205, 202, 201, 150],
            backgroundColor: "#52CDFF",
            borderColor: [
                '#52CDFF',
            ],
            borderWidth: 0,
            fill: true, // 3: no fill
              
          },{
            label: 'This week',
            data: [215, 290, 210, 250, 290, 230, 290, 210, 280, 220, 190, 300],
            backgroundColor: "#1F3BB3",
            borderColor: [
                '#1F3BB3',
            ],
            borderWidth: 0,
            fill: true, // 3: no fill
        }]
      };
  
      var marketingOverviewOptions = {
        responsive: true,
        maintainAspectRatio: false,
        scales: {
          yAxes: [{
              gridLines: {
                  display: true,
                  drawBorder: false,
                  color:"#F0F0F0",
                  zeroLineColor: '#F0F0F0',
              },
              ticks: {
                beginAtZero: true,
                autoSkip: true,
                maxTicksLimit: 5,
                fontSize: 10,
                color:"#6B778C"
              }
          }],
          xAxes: [{
            stacked: true,
            barPercentage: 0.35,
            gridLines: {
                display: false,
                drawBorder: false,
            },
            ticks: {
              beginAtZero: false,
              autoSkip: true,
              maxTicksLimit: 12,
              fontSize: 10,
              color:"#6B778C"
            }
        }],
        },
        legend:false,
        legendCallback: function (chart) {
          var text = [];
          text.push('<div class="chartjs-legend"><ul>');
          for (var i = 0; i < chart.data.datasets.length; i++) {
            console.log(chart.data.datasets[i]); // see what's inside the obj.
            text.push('<li class="text-muted text-small">');
            text.push('<span style="background-color:' + chart.data.datasets[i].borderColor + '">' + '</span>');
            text.push(chart.data.datasets[i].label);
            text.push('</li>');
          }
          text.push('</ul></div>');
          return text.join("");
        },
        
        elements: {
            line: {
                tension: 0.4,
            }
        },
        tooltips: {
          backgroundColor: 'rgba(31, 59, 179, 1)',
        }
      }
      var marketingOverview = new Chart(marketingOverviewChart, {
          type: 'bar',
          data: marketingOverviewData,
          options: marketingOverviewOptions
      });
      document.getElementById('marketing-overview-legend').innerHTML = marketingOverview.generateLegend();
    }
    if ($("#marketingOverview-dark").length) {
      var marketingOverviewChartDark = document.getElementById("marketingOverview-dark").getContext('2d');
      var marketingOverviewDataDark = {
          labels: ["JAN","FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"],
          datasets: [{
              label: 'Last week',
              data: [110, 220, 200, 190, 220, 110, 210, 110, 205, 202, 201, 150],
              backgroundColor: "#F29F67",
              borderColor: [
                  '#F29F67',
              ],
              borderWidth: 0,
              fill: true, // 3: no fill
              
          },{
            label: 'This week',
            data: [215, 290, 210, 250, 290, 230, 290, 210, 280, 220, 190, 300],
            backgroundColor: "#5A5B6A",
            borderColor: [
                '#5A5B6A',
            ],
            borderWidth: 0,
            fill: true, // 3: no fill
        }]
      };
  
      var marketingOverviewOptionsDark = {
        responsive: true,
        maintainAspectRatio: false,
          scales: {
              yAxes: [{
                  gridLines: {
                      display: true,
                      drawBorder: false,
                      color:"rgba(255,255,255,.05)",
                      zeroLineColor: "rgba(255,255,255,.05)",
                  },
                  ticks: {
                    beginAtZero: true,
                    autoSkip: true,
                    maxTicksLimit: 5,
                    fontSize: 10,
                    color:"#6B778C"
                  }
              }],
              xAxes: [{
                stacked: true,
                barPercentage: 0.35,
                gridLines: {
                    display: false,
                    drawBorder: false,
                },
                ticks: {
                  beginAtZero: false,
                  autoSkip: true,
                  maxTicksLimit: 7,
                  fontSize: 10,
                  color:"#6B778C"
                }
            }],
          },
          legend:false,
          legendCallback: function (chart) {
            var text = [];
            text.push('<div class="chartjs-legend"><ul>');
            for (var i = 0; i < chart.data.datasets.length; i++) {
              console.log(chart.data.datasets[i]); // see what's inside the obj.
              text.push('<li class="text-muted text-small">');
              text.push('<span style="background-color:' + chart.data.datasets[i].borderColor + '">' + '</span>');
              text.push(chart.data.datasets[i].label);
              text.push('</li>');
            }
            text.push('</ul></div>');
            return text.join("");
          },
          
          elements: {
              line: {
                  tension: 0.4,
              }
          },
          tooltips: {
              backgroundColor: 'rgba(31, 59, 179, 1)',
          }
      }
      var marketingOverviewDark = new Chart(marketingOverviewChartDark, {
          type: 'bar',
          data: marketingOverviewDataDark,
          options: marketingOverviewOptionsDark
      });
      document.getElementById('marketing-overview-legend').innerHTML = marketingOverviewDark.generateLegend();
    }
    if ($("#marketingOverviewBrown").length) {
      var marketingOverviewChart = document.getElementById("marketingOverviewBrown").getContext('2d');
      var marketingOverviewData = {
          labels: ["JAN","FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"],
          datasets: [{
              label: 'Last week',
              data: [110, 220, 200, 190, 220, 110, 210, 110, 205, 202, 201, 150],
              backgroundColor: "#F29F67",
              borderColor: [
                  '#F29F67',
              ],
              borderWidth: 0,
              fill: true, // 3: no fill
              
          },{
            label: 'This week',
            data: [215, 290, 210, 250, 290, 230, 290, 210, 280, 220, 190, 300],
            backgroundColor: "#1E1E2C",
            borderColor: [
                '#1E1E2C',
            ],
            borderWidth: 0,
            fill: true, // 3: no fill
        }]
      };
  
      var marketingOverviewOptions = {
        responsive: true,
        maintainAspectRatio: false,
          scales: {
              yAxes: [{
                  gridLines: {
                      display: true,
                      drawBorder: false,
                      color:"#F0F0F0",
                      zeroLineColor: '#F0F0F0',
                  },
                  ticks: {
                    beginAtZero: true,
                    autoSkip: true,
                    maxTicksLimit: 5,
                    fontSize: 10,
                    color:"#6B778C"
                  }
              }],
              xAxes: [{
                stacked: true,
                barPercentage: 0.35,
                gridLines: {
                    display: false,
                    drawBorder: false,
                },
                ticks: {
                  beginAtZero: false,
                  autoSkip: true,
                  maxTicksLimit: 12,
                  fontSize: 10,
                  color:"#6B778C"
                }
            }],
          },
          legend:false,
          legendCallback: function (chart) {
            var text = [];
            text.push('<div class="chartjs-legend"><ul>');
            for (var i = 0; i < chart.data.datasets.length; i++) {
              console.log(chart.data.datasets[i]); // see what's inside the obj.
              text.push('<li class="text-muted text-small">');
              text.push('<span style="background-color:' + chart.data.datasets[i].borderColor + '">' + '</span>');
              text.push(chart.data.datasets[i].label);
              text.push('</li>');
            }
            text.push('</ul></div>');
            return text.join("");
          },
          
          elements: {
              line: {
                  tension: 0.4,
              }
          },
          tooltips: {
              backgroundColor: 'rgba(31, 59, 179, 1)',
          }
      }
      var marketingOverview = new Chart(marketingOverviewChart, {
          type: 'bar',
          data: marketingOverviewData,
          options: marketingOverviewOptions
      });
      document.getElementById('marketing-overview-legend').innerHTML = marketingOverview.generateLegend();
    }
    if ($("#marketingOverviewPurple").length) {
      var marketingOverviewChart = document.getElementById("marketingOverviewPurple").getContext('2d');
      var marketingOverviewData = {
          labels: ["JAN","FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"],
          datasets: [{
            label: 'Last week',
            data: [110, 220, 200, 190, 220, 110, 210, 110, 205, 202, 201, 150],
            backgroundColor: "#CACEEA",
            borderColor: [
                '#CACEEA',
            ],
            borderWidth: 0,
            fill: true, // 3: no fill
              
          },{
            label: 'This week',
            data: [215, 290, 210, 250, 290, 230, 290, 210, 280, 220, 190, 300],
            backgroundColor: "#7B20C7",
            borderColor: [
                '#7B20C7',
            ],
            borderWidth: 0,
            fill: true, // 3: no fill
        }]
      };
  
      var marketingOverviewOptions = {
        responsive: true,
        maintainAspectRatio: false,
        scales: {
          yAxes: [{
              gridLines: {
                  display: true,
                  drawBorder: false,
                  color:"#F0F0F0",
                  zeroLineColor: '#F0F0F0',
              },
              ticks: {
                beginAtZero: true,
                autoSkip: true,
                maxTicksLimit: 5,
                fontSize: 10,
                color:"#6B778C"
              }
          }],
          xAxes: [{
            stacked: true,
            barPercentage: 0.35,
            gridLines: {
                display: false,
                drawBorder: false,
            },
            ticks: {
              beginAtZero: false,
              autoSkip: true,
              maxTicksLimit: 12,
              fontSize: 10,
              color:"#6B778C"
            }
        }],
        },
        legend:false,
        legendCallback: function (chart) {
          var text = [];
          text.push('<div class="chartjs-legend"><ul>');
          for (var i = 0; i < chart.data.datasets.length; i++) {
            console.log(chart.data.datasets[i]); // see what's inside the obj.
            text.push('<li class="text-muted text-small">');
            text.push('<span style="background-color:' + chart.data.datasets[i].borderColor + '">' + '</span>');
            text.push(chart.data.datasets[i].label);
            text.push('</li>');
          }
          text.push('</ul></div>');
          return text.join("");
        },
        
        elements: {
            line: {
                tension: 0.4,
            }
        },
        tooltips: {
          backgroundColor: 'rgba(31, 59, 179, 1)',
        }
      }
      var marketingOverview = new Chart(marketingOverviewChart, {
          type: 'bar',
          data: marketingOverviewData,
          options: marketingOverviewOptions
      });
      document.getElementById('marketing-overview-legend').innerHTML = marketingOverview.generateLegend();
    }
    if ($("#marketingOverviewPurple-dark").length) {
      var marketingOverviewChart = document.getElementById("marketingOverviewPurple-dark").getContext('2d');
      var marketingOverviewData = {
          labels: ["JAN","FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"],
          datasets: [{
            label: 'Last week',
            data: [110, 220, 200, 190, 220, 110, 210, 110, 205, 202, 201, 150],
            backgroundColor: "#EBD3FF",
            borderColor: [
                '#EBD3FF',
            ],
            borderWidth: 0,
            fill: true, // 3: no fill
              
          },{
            label: 'This week',
            data: [215, 290, 210, 250, 290, 230, 290, 210, 280, 220, 190, 300],
            backgroundColor: "#BE70FF",
            borderColor: [
                '#BE70FF',
            ],
            borderWidth: 0,
            fill: true, // 3: no fill
        }]
      };
  
      var marketingOverviewOptions = {
        responsive: true,
        maintainAspectRatio: false,
        scales: {
          yAxes: [{
              gridLines: {
                  display: true,
                  drawBorder: false,
                  color:"rgba(232,232,232,.12)",
                  zeroLineColor: 'rgba(232,232,232,.12)',
              },
              ticks: {
                beginAtZero: true,
                autoSkip: true,
                maxTicksLimit: 5,
                fontSize: 10,
                color:"rgba(232,232,232,.12)"
              }
          }],
          xAxes: [{
            stacked: true,
            barPercentage: 0.35,
            gridLines: {
                display: false,
                drawBorder: false,
            },
            ticks: {
              beginAtZero: false,
              autoSkip: true,
              maxTicksLimit: 12,
              fontSize: 10,
              color:"rgba(232,232,232,.12)"
            }
        }],
        },
        legend:false,
        legendCallback: function (chart) {
          var text = [];
          text.push('<div class="chartjs-legend"><ul>');
          for (var i = 0; i < chart.data.datasets.length; i++) {
            console.log(chart.data.datasets[i]); // see what's inside the obj.
            text.push('<li class="text-muted text-small">');
            text.push('<span style="background-color:' + chart.data.datasets[i].borderColor + '">' + '</span>');
            text.push(chart.data.datasets[i].label);
            text.push('</li>');
          }
          text.push('</ul></div>');
          return text.join("");
        },
        
        elements: {
            line: {
                tension: 0.4,
            }
        },
        tooltips: {
          backgroundColor: 'rgba(31, 59, 179, 1)',
        }
      }
      var marketingOverview = new Chart(marketingOverviewChart, {
          type: 'bar',
          data: marketingOverviewData,
          options: marketingOverviewOptions
      });
      document.getElementById('marketing-overview-legend').innerHTML = marketingOverview.generateLegend();
    }
    if ($("#marketingOverviewCrm").length) {
      var marketingOverviewChart = document.getElementById("marketingOverviewCrm").getContext('2d');
      var marketingOverviewData = {
          labels: ["Mon","Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
          datasets: [{
            label: 'Last week',
            data: [350, 500, 100, 400, 550, 310, 240],
            backgroundColor: "#2A21BA",
            borderColor: [
                '#2A21BA',
            ],
            borderWidth: 0,
            fill: true, // 3: no fill
              
          }]
      };
  
      var marketingOverviewOptions = {
        responsive: true,
        maintainAspectRatio: false,
        scales: {
          yAxes: [{
              gridLines: {
                  display: true,
                  drawBorder: false,
                  color:"#F0F0F0",
                  zeroLineColor: '#F0F0F0',
              },
              ticks: {
                beginAtZero: true,
                autoSkip: true,
                maxTicksLimit: 6,
                fontSize: 10,
                color:"#6B778C"
              }
          }],
          xAxes: [{
            stacked: true,
            barPercentage: 0.5,
            gridLines: {
                display: false,
                drawBorder: false,
            },
            ticks: {
              beginAtZero: false,
              autoSkip: true,
              maxTicksLimit: 12,
              fontSize: 10,
              color:"#6B778C"
            }
        }],
        },
        legend:false,
        
        elements: {
            line: {
                tension: 0.4,
            }
        },
        tooltips: {
          backgroundColor: 'rgba(31, 59, 179, 1)',
        }
      }
      var marketingOverview = new Chart(marketingOverviewChart, {
          type: 'bar',
          data: marketingOverviewData,
          options: marketingOverviewOptions
      });
    }
    if ($("#doughnutChart").length) {
      var doughnutChartCanvas = $("#doughnutChart").get(0).getContext("2d");
      var doughnutPieData = {
        datasets: [{
          data: [40, 20, 30, 10],
          backgroundColor: [
            "#1F3BB3",
            "#FDD0C7",
            "#52CDFF",
            "#81DADA"
          ],
          borderColor: [
            "#1F3BB3",
            "#FDD0C7",
            "#52CDFF",
            "#81DADA"
          ],
        }],
  
        // These labels appear in the legend and in the tooltips when hovering different arcs
        labels: [
          'Total',
          'Net',
          'Gross',
          'AVG',
        ]
      };
      var doughnutPieOptions = {
        cutoutPercentage: 50,
        animationEasing: "easeOutBounce",
        animateRotate: true,
        animateScale: false,
        responsive: true,
        maintainAspectRatio: true,
        showScale: true,
        legend: false,
        legendCallback: function (chart) {
          var text = [];
          text.push('<div class="chartjs-legend"><ul class="justify-content-center">');
          for (var i = 0; i < chart.data.datasets[0].data.length; i++) {
            text.push('<li><span style="background-color:' + chart.data.datasets[0].backgroundColor[i] + '">');
            text.push('</span>');
            if (chart.data.labels[i]) {
              text.push(chart.data.labels[i]);
            }
            text.push('</li>');
          }
          text.push('</div></ul>');
          return text.join("");
        },
        
        layout: {
          padding: {
            left: 0,
            right: 0,
            top: 0,
            bottom: 0
          }
        },
        tooltips: {
          callbacks: {
            title: function(tooltipItem, data) {
              return data['labels'][tooltipItem[0]['index']];
            },
            label: function(tooltipItem, data) {
              return data['datasets'][0]['data'][tooltipItem['index']];
            }
          },
            
          backgroundColor: '#fff',
          titleFontSize: 14,
          titleFontColor: '#0B0F32',
          bodyFontColor: '#737F8B',
          bodyFontSize: 11,
          displayColors: false
        }
      };
      var doughnutChart = new Chart(doughnutChartCanvas, {
        type: 'doughnut',
        data: doughnutPieData,
        options: doughnutPieOptions
      });
      document.getElementById('doughnut-chart-legend').innerHTML = doughnutChart.generateLegend();
    }
    if ($("#doughnutChartBrown").length) {
      var doughnutChartCanvas = $("#doughnutChartBrown").get(0).getContext("2d");
      var doughnutPieData = {
        datasets: [{
          data: [40, 20, 30, 10],
          backgroundColor: [
            "#1E1E2C",
            "#F3C5BE",
            "#75CDCD",
            "#F29F67"
          ],
          borderColor: [
            "#1E1E2C",
            "#F3C5BE",
            "#75CDCD",
            "#F29F67"
          ],
        }],
  
        // These labels appear in the legend and in the tooltips when hovering different arcs
        labels: [
          'Total',
          'Net',
          'Gross',
          'AVG',
        ]
      };
      var doughnutPieOptions = {
        cutoutPercentage: 50,
        animationEasing: "easeOutBounce",
        animateRotate: true,
        animateScale: false,
        responsive: true,
        maintainAspectRatio: true,
        showScale: true,
        legend: false,
        legendCallback: function (chart) {
          var text = [];
          text.push('<div class="chartjs-legend"><ul class="justify-content-center">');
          for (var i = 0; i < chart.data.datasets[0].data.length; i++) {
            text.push('<li><span style="background-color:' + chart.data.datasets[0].backgroundColor[i] + '">');
            text.push('</span>');
            if (chart.data.labels[i]) {
              text.push(chart.data.labels[i]);
            }
            text.push('</li>');
          }
          text.push('</div></ul>');
          return text.join("");
        },
        
        layout: {
          padding: {
            left: 0,
            right: 0,
            top: 0,
            bottom: 0
          }
        },
        tooltips: {
          callbacks: {
            title: function(tooltipItem, data) {
              return data['labels'][tooltipItem[0]['index']];
            },
            label: function(tooltipItem, data) {
              return data['datasets'][0]['data'][tooltipItem['index']];
            }
          },
            
          backgroundColor: '#fff',
          titleFontSize: 14,
          titleFontColor: '#0B0F32',
          bodyFontColor: '#737F8B',
          bodyFontSize: 11,
          displayColors: false
        }
      };
      var doughnutChart = new Chart(doughnutChartCanvas, {
        type: 'doughnut',
        data: doughnutPieData,
        options: doughnutPieOptions
      });
      document.getElementById('doughnut-chart-legend').innerHTML = doughnutChart.generateLegend();
    }
    if ($("#doughnutChartDark").length) {
      var doughnutChartDarkCanvas = $("#doughnutChartDark").get(0).getContext("2d");
      var doughnutPieDarkData = {
        datasets: [{
          data: [40, 20, 30, 10],
          backgroundColor: [
            "#2A4B7A",
            "#F3C5BE",
            "#75CDCD",
            "#F29F67"
          ],
          borderColor: [
            "#2A4B7A",
            "#F3C5BE",
            "#75CDCD",
            "#F29F67"
          ],
        }],
  
        // These labels appear in the legend and in the tooltips when hovering different arcs
        labels: [
          'Total',
          'Net',
          'Gross',
          'AVG',
        ]
      };
      var doughnutPieDarkOptions = {
        cutoutPercentage: 50,
        animationEasing: "easeOutBounce",
        animateRotate: true,
        animateScale: false,
        responsive: true,
        maintainAspectRatio: true,
        showScale: true,
        legend: false,
        legendCallback: function (chart) {
          var text = [];
          text.push('<div class="chartjs-legend"><ul class="justify-content-center">');
          for (var i = 0; i < chart.data.datasets[0].data.length; i++) {
            text.push('<li><span style="background-color:' + chart.data.datasets[0].backgroundColor[i] + '">');
            text.push('</span>');
            if (chart.data.labels[i]) {
              text.push(chart.data.labels[i]);
            }
            text.push('</li>');
          }
          text.push('</div></ul>');
          return text.join("");
        },
        
        layout: {
          padding: {
            left: 0,
            right: 0,
            top: 0,
            bottom: 0
          }
        },
        tooltips: {
          callbacks: {
            title: function(tooltipItem, data) {
              return data['labels'][tooltipItem[0]['index']];
            },
            label: function(tooltipItem, data) {
              return data['datasets'][0]['data'][tooltipItem['index']];
            }
          },
            
          backgroundColor: '#191B20',
          titleFontSize: 14,
          titleFontColor: '#D8D9E3',
          bodyFontColor: '#808191',
          bodyFontSize: 11,
          displayColors: false
        }
      };
      var doughnutChartDark = new Chart(doughnutChartDarkCanvas, {
        type: 'doughnut',
        data: doughnutPieDarkData,
        options: doughnutPieDarkOptions
      });
      document.getElementById('doughnut-chart-legend').innerHTML = doughnutChartDark.generateLegend();
    }
    if ($("#doughnutChartCrm").length) {
      var doughnutChartCanvas = $("#doughnutChartCrm").get(0).getContext("2d");
      var doughnutPieData = {
        datasets: [{
          data: [40, 30, 30],
          backgroundColor: [
            "#1F3BB3",
            "#00CDFF",
            "#00AAB6"
          ],
          borderColor: [
            "#fff",
            "#fff",
            "#fff"
          ],
        }],
  
        // These labels appear in the legend and in the tooltips when hovering different arcs
        labels: [
          'Branch 1  ( 30% )',
          'Branch 2  ( 40% )',
          'Branch 3  ( 30% )'
        ]
      };
      var doughnutPieOptions = {
        cutoutPercentage: 60,
        animationEasing: "easeOutBounce",
        animateRotate: true,
        animateScale: false,
        responsive: true,
        maintainAspectRatio: true,
        showScale: true,
        legend: false,
        legendCallback: function (chart) {
          var text = [];
          text.push('<div class="chartjs-legend"><ul class="justify-content-center d-block">');
          for (var i = 0; i < chart.data.datasets[0].data.length; i++) {
            text.push('<li class="d-block text-medium mb-3"><span style="background-color:' + chart.data.datasets[0].backgroundColor[i] + '">');
            text.push('</span>');
            if (chart.data.labels[i]) {
              text.push(chart.data.labels[i]);
            }
            text.push('</li>');
          }
          text.push('</div></ul>');
          return text.join("");
        },
        
        layout: {
          padding: {
            left: 0,
            right: 0,
            top: 0,
            bottom: 0
          }
        },
        tooltips: {
          callbacks: {
            title: function(tooltipItem, data) {
              return data['labels'][tooltipItem[0]['index']];
            },
            label: function(tooltipItem, data) {
              return data['datasets'][0]['data'][tooltipItem['index']];
            }
          },
            
          backgroundColor: '#fff',
          titleFontSize: 14,
          titleFontColor: '#0B0F32',
          bodyFontColor: '#737F8B',
          bodyFontSize: 11,
          displayColors: false
        }
      };
      var doughnutChart = new Chart(doughnutChartCanvas, {
        type: 'doughnut',
        data: doughnutPieData,
        options: doughnutPieOptions
      });
      document.getElementById('doughnut-chart-legend-crm').innerHTML = doughnutChart.generateLegend();
    }
    if ($("#doughnutChartPurple").length) {
      var doughnutChartCanvas = $("#doughnutChartPurple").get(0).getContext("2d");
      var doughnutPieData = {
        datasets: [{
          data: [40, 20, 30, 10],
          backgroundColor: [
            "#7B20C7",
            "#006CFF",
            "#00CCCC",
            "#ADB2BD"
          ],
          borderColor: [
            "#fff",
            "#fff",
            "#fff",
            "#fff"
          ],
        }],
  
        // These labels appear in the legend and in the tooltips when hovering different arcs
        labels: [
          'Total',
          'Net',
          'Gross',
          'AVG',
        ]
      };
      var doughnutPieOptions = {
        cutoutPercentage: 50,
        animationEasing: "easeOutBounce",
        animateRotate: true,
        animateScale: false,
        responsive: true,
        maintainAspectRatio: true,
        showScale: true,
        legend: false,
        legendCallback: function (chart) {
          var text = [];
          text.push('<div class="chartjs-legend"><ul class="justify-content-center">');
          for (var i = 0; i < chart.data.datasets[0].data.length; i++) {
            text.push('<li><span style="background-color:' + chart.data.datasets[0].backgroundColor[i] + '">');
            text.push('</span>');
            if (chart.data.labels[i]) {
              text.push(chart.data.labels[i]);
            }
            text.push('</li>');
          }
          text.push('</div></ul>');
          return text.join("");
        },
        
        layout: {
          padding: {
            left: 0,
            right: 0,
            top: 0,
            bottom: 0
          }
        },
        tooltips: {
          callbacks: {
            title: function(tooltipItem, data) {
              return data['labels'][tooltipItem[0]['index']];
            },
            label: function(tooltipItem, data) {
              return data['datasets'][0]['data'][tooltipItem['index']];
            }
          },
            
          backgroundColor: '#fff',
          titleFontSize: 14,
          titleFontColor: '#0B0F32',
          bodyFontColor: '#737F8B',
          bodyFontSize: 11,
          displayColors: false
        }
      };
      var doughnutChart = new Chart(doughnutChartCanvas, {
        type: 'doughnut',
        data: doughnutPieData,
        options: doughnutPieOptions
      });
      document.getElementById('doughnut-chart-legend').innerHTML = doughnutChart.generateLegend();
    }
    if ($("#doughnutChartPurple-dark").length) {
      var doughnutChartCanvas = $("#doughnutChartPurple-dark").get(0).getContext("2d");
      var doughnutPieData = {
        datasets: [{
          data: [40, 20, 30, 10],
          backgroundColor: [
            "#BE70FF",
            "#00A3FF",
            "#00CCCC",
            "#ADB2BD"
          ],
          borderColor: [
            "#000",
            "#000",
            "#000",
            "#000"
          ],
        }],
  
        // These labels appear in the legend and in the tooltips when hovering different arcs
        labels: [
          'Total',
          'Net',
          'Gross',
          'AVG',
        ]
      };
      var doughnutPieOptions = {
        cutoutPercentage: 50,
        animationEasing: "easeOutBounce",
        animateRotate: true,
        animateScale: false,
        responsive: true,
        maintainAspectRatio: true,
        showScale: true,
        legend: false,
        legendCallback: function (chart) {
          var text = [];
          text.push('<div class="chartjs-legend"><ul class="justify-content-center">');
          for (var i = 0; i < chart.data.datasets[0].data.length; i++) {
            text.push('<li><span style="background-color:' + chart.data.datasets[0].backgroundColor[i] + '">');
            text.push('</span>');
            if (chart.data.labels[i]) {
              text.push(chart.data.labels[i]);
            }
            text.push('</li>');
          }
          text.push('</div></ul>');
          return text.join("");
        },
        
        layout: {
          padding: {
            left: 0,
            right: 0,
            top: 0,
            bottom: 0
          }
        },
        tooltips: {
          callbacks: {
            title: function(tooltipItem, data) {
              return data['labels'][tooltipItem[0]['index']];
            },
            label: function(tooltipItem, data) {
              return data['datasets'][0]['data'][tooltipItem['index']];
            }
          },
            
          backgroundColor: '#fff',
          titleFontSize: 14,
          titleFontColor: '#0B0F32',
          bodyFontColor: '#737F8B',
          bodyFontSize: 11,
          displayColors: false
        }
      };
      var doughnutChart = new Chart(doughnutChartCanvas, {
        type: 'doughnut',
        data: doughnutPieData,
        options: doughnutPieOptions
      });
      document.getElementById('doughnut-chart-legend').innerHTML = doughnutChart.generateLegend();
    }
    if ($("#doughnutChartSales").length) {
      var doughnutChartCanvas = $("#doughnutChartSales").get(0).getContext("2d");
      var doughnutPieData = {
        datasets: [{
          data: [40, 30, 30],
          backgroundColor: [
            "#1F3BB3",
            "#00CDFF",
            "#00AAB6"
          ],
          borderColor: [
            "#fff",
            "#fff",
            "#fff"
          ],
        }],
  
      };
      var doughnutPieOptions = {
        cutoutPercentage: 60,
        animationEasing: "easeOutBounce",
        animateRotate: true,
        animateScale: false,
        responsive: true,
        maintainAspectRatio: true,
        showScale: true,
        legend: false,
        
        layout: {
          padding: {
            left: 0,
            right: 0,
            top: 0,
            bottom: 0
          }
        },
        tooltips: {
            
          backgroundColor: '#fff',
          titleFontSize: 14,
          titleFontColor: '#0B0F32',
          bodyFontColor: '#737F8B',
          bodyFontSize: 11,
          displayColors: false
        }
      };
      var doughnutChart = new Chart(doughnutChartCanvas, {
        type: 'doughnut',
        data: doughnutPieData,
        options: doughnutPieOptions
      });
    }
    if ($("#leaveReport").length) {
      var leaveReportChart = document.getElementById("leaveReport").getContext('2d');
      var leaveReportData = {
          labels: ["Jan","Feb", "Mar", "Apr", "May"],
          datasets: [{
              label: 'Last week',
              data: [18, 25, 39, 11, 24],
              backgroundColor: "#52CDFF",
              borderColor: [
                  '#52CDFF',
              ],
              borderWidth: 0,
              fill: true, // 3: no fill
              
          }]
      };
  
      var leaveReportOptions = {
        responsive: true,
        maintainAspectRatio: false,
          scales: {
              yAxes: [{
                  gridLines: {
                      display: true,
                      drawBorder: false,
                      color:"rgba(255,255,255,.05)",
                      zeroLineColor: "rgba(255,255,255,.05)",
                  },
                  ticks: {
                    beginAtZero: true,
                    autoSkip: true,
                    maxTicksLimit: 5,
                    fontSize: 10,
                    color:"#6B778C"
                  }
              }],
              xAxes: [{
                barPercentage: 0.5,
                gridLines: {
                    display: false,
                    drawBorder: false,
                },
                ticks: {
                  beginAtZero: false,
                  autoSkip: true,
                  maxTicksLimit: 7,
                  fontSize: 10,
                  color:"#6B778C"
                }
            }],
          },
          legend:false,
          
          elements: {
              line: {
                  tension: 0.4,
              }
          },
          tooltips: {
              backgroundColor: 'rgba(31, 59, 179, 1)',
          }
      }
      var leaveReport = new Chart(leaveReportChart, {
          type: 'bar',
          data: leaveReportData,
          options: leaveReportOptions
      });
    }
    if ($("#leaveReport-dark").length) {
      var leaveReportChartDark = document.getElementById("leaveReport-dark").getContext('2d');
      var leaveReportDataDark = {
          labels: ["JAN","FEB", "MAR", "APR", "MAY"],
          datasets: [{
              label: 'Last week',
              data: [18, 25, 39, 11, 24],
              backgroundColor: "#F29F67",
              borderColor: [
                  '#F29F67',
              ],
              borderWidth: 0,
              fill: true, // 3: no fill
              
          }]
      };
  
      var leaveReportOptionsDark = {
        responsive: true,
        maintainAspectRatio: false,
          scales: {
              yAxes: [{
                  gridLines: {
                      display: true,
                      drawBorder: false,
                      color:"#383e5d",
                      zeroLineColor: '#383e5d',
                  },
                  ticks: {
                    beginAtZero: true,
                    autoSkip: true,
                    maxTicksLimit: 5,
                    fontSize: 10,
                    color:"#6B778C"
                  }
              }],
              xAxes: [{
                barPercentage: 0.5,
                gridLines: {
                    display: false,
                    drawBorder: false,
                },
                ticks: {
                  beginAtZero: false,
                  autoSkip: true,
                  maxTicksLimit: 7,
                  fontSize: 10,
                  color:"#6B778C"
                }
            }],
          },
          legend:false,
          
          elements: {
              line: {
                  tension: 0.4,
              }
          },
          tooltips: {
              backgroundColor: 'rgba(31, 59, 179, 1)',
          }
      }
      var leaveReportDark = new Chart(leaveReportChartDark, {
          type: 'bar',
          data: leaveReportDataDark,
          options: leaveReportOptionsDark
      });
    }
    if ($("#leaveReportBrown").length) {
      var leaveReportChart = document.getElementById("leaveReportBrown").getContext('2d');
      var leaveReportData = {
          labels: ["Jan","Feb", "Mar", "Apr", "May"],
          datasets: [{
              label: 'Last week',
              data: [18, 25, 39, 11, 24],
              backgroundColor: "#F29F67",
              borderColor: [
                  '#F29F67',
              ],
              borderWidth: 0,
              fill: true, // 3: no fill
              
          }]
      };
  
      var leaveReportOptions = {
        responsive: true,
        maintainAspectRatio: false,
          scales: {
              yAxes: [{
                  gridLines: {
                      display: true,
                      drawBorder: false,
                      color:"rgba(255,255,255,.05)",
                      zeroLineColor: "rgba(255,255,255,.05)",
                  },
                  ticks: {
                    beginAtZero: true,
                    autoSkip: true,
                    maxTicksLimit: 5,
                    fontSize: 10,
                    color:"#6B778C"
                  }
              }],
              xAxes: [{
                barPercentage: 0.5,
                gridLines: {
                    display: false,
                    drawBorder: false,
                },
                ticks: {
                  beginAtZero: false,
                  autoSkip: true,
                  maxTicksLimit: 7,
                  fontSize: 10,
                  color:"#6B778C"
                }
            }],
          },
          legend:false,
          
          elements: {
              line: {
                  tension: 0.4,
              }
          },
          tooltips: {
              backgroundColor: 'rgba(31, 59, 179, 1)',
          }
      }
      var leaveReport = new Chart(leaveReportChart, {
          type: 'bar',
          data: leaveReportData,
          options: leaveReportOptions
      });
    }
    if ($("#leaveReportPurple").length) {
      var leaveReportChart = document.getElementById("leaveReportPurple").getContext('2d');
      var leaveReportData = {
          labels: ["Jan","Feb", "Mar", "Apr", "May"],
          datasets: [{
              label: 'Last week',
              data: [18, 25, 39, 11, 24],
              backgroundColor: "#006CFF",
              borderColor: [
                  '#006CFF',
              ],
              borderWidth: 0,
              fill: true, // 3: no fill
              
          }]
      };
  
      var leaveReportOptions = {
        responsive: true,
        maintainAspectRatio: false,
          scales: {
              yAxes: [{
                  gridLines: {
                      display: true,
                      drawBorder: false,
                      color:"rgba(255,255,255,.05)",
                      zeroLineColor: "rgba(255,255,255,.05)",
                  },
                  ticks: {
                    beginAtZero: true,
                    autoSkip: true,
                    maxTicksLimit: 5,
                    fontSize: 10,
                    color:"#6B778C"
                  }
              }],
              xAxes: [{
                barPercentage: 0.5,
                gridLines: {
                    display: false,
                    drawBorder: false,
                },
                ticks: {
                  beginAtZero: false,
                  autoSkip: true,
                  maxTicksLimit: 7,
                  fontSize: 10,
                  color:"#6B778C"
                }
            }],
          },
          legend:false,
          
          elements: {
              line: {
                  tension: 0.4,
              }
          },
          tooltips: {
              backgroundColor: 'rgba(31, 59, 179, 1)',
          }
      }
      var leaveReport = new Chart(leaveReportChart, {
          type: 'bar',
          data: leaveReportData,
          options: leaveReportOptions
      });
    }
    if ($("#leaveReportPurple-dark").length) {
      var leaveReportChart = document.getElementById("leaveReportPurple-dark").getContext('2d');
      var leaveReportData = {
          labels: ["Jan","Feb", "Mar", "Apr", "May"],
          datasets: [{
              label: 'Last week',
              data: [18, 25, 39, 11, 24],
              backgroundColor: "#00A3FF",
              borderColor: [
                  '#00A3FF',
              ],
              borderWidth: 0,
              fill: true, // 3: no fill
              
          }]
      };
  
      var leaveReportOptions = {
        responsive: true,
        maintainAspectRatio: false,
          scales: {
              yAxes: [{
                  gridLines: {
                      display: true,
                      drawBorder: false,
                      color:"rgba(255,255,255,.05)",
                      zeroLineColor: "rgba(255,255,255,.05)",
                  },
                  ticks: {
                    beginAtZero: true,
                    autoSkip: true,
                    maxTicksLimit: 5,
                    fontSize: 10,
                    color:"#6B778C"
                  }
              }],
              xAxes: [{
                barPercentage: 0.5,
                gridLines: {
                    display: false,
                    drawBorder: false,
                },
                ticks: {
                  beginAtZero: false,
                  autoSkip: true,
                  maxTicksLimit: 7,
                  fontSize: 10,
                  color:"#6B778C"
                }
            }],
          },
          legend:false,
          
          elements: {
              line: {
                  tension: 0.4,
              }
          },
          tooltips: {
              backgroundColor: 'rgba(31, 59, 179, 1)',
          }
      }
      var leaveReport = new Chart(leaveReportChart, {
          type: 'bar',
          data: leaveReportData,
          options: leaveReportOptions
      });
    }
    if ($("#salesReportSales").length) {
      var marketingOverviewChart = document.getElementById("salesReportSales").getContext('2d');
      
      var marketingOverviewData = {
          labels: ["Mon","Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
          datasets: [{
            label: 'Last week',
            data: [350, 500, 100, 400, 550, 310, 240],
            backgroundColor: ["#F95F53", "#00CDFF", "#F95F53", "#00CDFF", "#00CDFF", "#F95F53", "#00CDFF"],
            borderColor: ["#F95F53", "#00CDFF", "#F95F53", "#00CDFF", "#00CDFF", "#F95F53", "#00CDFF"],
            borderWidth: 0,
            borderRadius: 50,
            fill: true, // 3: no fill
              
          }]
      };
  
      var marketingOverviewOptions = {
        responsive: true,
        maintainAspectRatio: false,
        scales: {
          yAxes: [{
              gridLines: {
                  display: true,
                  drawBorder: false,
                  color:"#F0F0F0",
                  zeroLineColor: '#F0F0F0',
              },
              ticks: {
                beginAtZero: true,
                autoSkip: true,
                maxTicksLimit: 6,
                fontSize: 10,
                color:"#6B778C"
              }
          }],
          xAxes: [{
            stacked: true,
            barPercentage: 0.5,
            gridLines: {
                display: false,
                drawBorder: false,
            },
            ticks: {
              beginAtZero: false,
              autoSkip: true,
              maxTicksLimit: 12,
              fontSize: 10,
              color:"#6B778C"
            }
        }],
        },
        legend:false,
        
        elements: {
            line: {
                tension: 0.4,
            }
        },
        tooltips: {
          backgroundColor: 'rgba(31, 59, 179, 1)',
        }
      }
      var marketingOverview = new Chart(marketingOverviewChart, {
          type: 'bar',
          data: marketingOverviewData,
          options: marketingOverviewOptions
      });
    }
    if ($('#totalFollowers').length) {

      var bar = new ProgressBar.Circle(totalFollowers, {
        color: '#000',
        svgStyle: {
          strokeLinecap: 'round',
        },
        // This has to be the same size as the maximum width to
        // prevent clipping
        strokeWidth: 20,
        trailWidth: 20,
        easing: 'easeInOut',
        duration: 1400,
        text: { 
          autoStyleContainer: false
        },
        from: {
          color: '#203BB3',
          width: 20,
          radius: 100,
        },
        to: {
          color: '#203BB3',
          width: 20
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
  
      bar.text.style.fontSize = '1.5rem';
      bar.text.style.fontWeight = 'bold';
      bar.animate(.80); // Number from 0.0 to 1.0
    }
    if ($('#totalCampaigns').length) {
      var bar = new ProgressBar.Circle(totalCampaigns, {
        color: '#000',
        svgStyle: {
          strokeLinecap: 'round',
        },
        // This has to be the same size as the maximum width to
        // prevent clipping
        strokeWidth: 20,
        trailWidth: 20,
        easing: 'easeInOut',
        duration: 1400,
        text: { 
          autoStyleContainer: false
        },
        from: {
          color: '#00CDFF',
          width: 20,
          radius: 100,
        },
        to: {
          color: '#00CDFF',
          width: 20
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
  
      bar.text.style.fontSize = '1.5rem';
      bar.text.style.fontWeight = 'bold';
      bar.animate(.80); // Number from 0.0 to 1.0
    }
    if ($("#salesTrendSales").length) {
      var graphGradient = document.getElementById("salesTrendSales").getContext('2d');
      var graphGradient2 = document.getElementById("salesTrendSales").getContext('2d');
      var saleGradientBg = graphGradient.createLinearGradient(5, 0, 5, 100);
      saleGradientBg.addColorStop(0, 'rgba(42, 33, 186, 0.2)');
      saleGradientBg.addColorStop(1, 'rgba(42, 33, 186, 0)');
      var saleGradientBg2 = graphGradient2.createLinearGradient(100, 0, 50, 150);
      saleGradientBg2.addColorStop(0, 'rgba(0, 205, 255, 0.2)');
      saleGradientBg2.addColorStop(1, 'rgba(0, 205, 255, 0)');
      var salesTopData = {
          labels: ["SUN","sun", "MON", "mon", "TUE","tue", "WED", "wed", "THU", "thu", "FRI", "fri", "SAT"],
          datasets: [{
              label: 'Online Payment',
              data: [50, 110, 60, 290, 200, 115, 130, 170, 90, 210, 240, 280, 200],
              backgroundColor: saleGradientBg,
              borderColor: [
                  '#2A21BA',
              ],
              borderWidth: 1.5,
              fill: true, // 3: no fill
              pointBorderWidth: 1,
              pointRadius: [0, 0, 0, 0, 0,0, 0, 0, 6, 0,0, 0, 0],
              pointHoverRadius: [0, 0, 0, 0, 0,0, 0, 0, 6, 0,0, 0, 0],
              pointBackgroundColor: ['','','','','','','','','#1F3BB3','','','','',''],
              pointBorderColor: ['','','','','','','','','#fff','','','','',''],
          },{
            label: 'Offline Sales',
            data: [30, 150, 190, 250, 120, 150, 130, 20, 30, 15, 40, 95, 180],
            backgroundColor: saleGradientBg2,
            borderColor: [
                '#52CDFF',
            ],
            borderWidth: 1.5,
            fill: true, // 3: no fill
            pointBorderWidth: 1,
            pointRadius: [0, 0, 0, 0, 0,0, 0, 0, 0, 0,0, 0, 0],
            pointHoverRadius: [0, 0, 0, 0, 0,0, 0, 0, 0, 0,0, 0, 0],
            pointBackgroundColor: ['','','','','','','','','','','','','',''],
            pointBorderColor: ['','','','','','','','','','','','','',''],
        }]
      };
  
      var salesTopOptions = {
        responsive: true,
        maintainAspectRatio: false,
          scales: {
              yAxes: [{
                  gridLines: {
                      display: true,
                      drawBorder: false,
                      color:"#F0F0F0",
                      zeroLineColor: '#F0F0F0',
                  },
                  ticks: {
                    beginAtZero: false,
                    autoSkip: true,
                    maxTicksLimit: 4,
                    fontSize: 10,
                    color:"#6B778C"
                  }
              }],
              xAxes: [{
                gridLines: {
                    display: false,
                    drawBorder: false,
                },
                ticks: {
                  beginAtZero: false,
                  autoSkip: true,
                  maxTicksLimit: 7,
                  fontSize: 10,
                  color:"#6B778C"
                }
            }],
          },
          legend:false,
          legendCallback: function (chart) {
            var text = [];
            text.push('<div class="chartjs-legend"><ul>');
            for (var i = 0; i < chart.data.datasets.length; i++) {
              console.log(chart.data.datasets[i]); // see what's inside the obj.
              text.push('<li>');
              text.push('<span class="legend-lg" style="background-color:' + chart.data.datasets[i].borderColor + '">' + '</span>');
              text.push(chart.data.datasets[i].label);
              text.push('</li>');
            }
            text.push('</ul></div>');
            return text.join("");
          },
          
          elements: {
              line: {
                  tension: 0.4,
              }
          },
          tooltips: {
              backgroundColor: 'rgba(31, 59, 179, 1)',
          }
      }
      var salesTop = new Chart(graphGradient, {
          type: 'line',
          data: salesTopData,
          options: salesTopOptions
      });
      document.getElementById('sales-trend-legend').innerHTML = salesTop.generateLegend();
    }
    $('[data-toggle="tooltip"]').tooltip(); 
    if ($('#workingFormats').length) {

      var bar = new ProgressBar.Circle(workingFormats, {
        color: '#000',
        svgStyle: {
          strokeLinecap: 'round',
        },
        // This has to be the same size as the maximum width to
        // prevent clipping
        strokeWidth: 10,
        trailWidth: 8,
        easing: 'easeInOut',
        duration: 1400,
        text: { 
          autoStyleContainer: false
        },
        from: {
          color: '#203BB3',
          width: 10,
          radius: 100,
        },
        to: {
          color: '#203BB3',
          width: 10
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
  
      bar.text.style.fontSize = '1.875rem';
      bar.text.style.fontWeight = 'bold';
      bar.animate(.30); // Number from 0.0 to 1.0
    }
    if ($("#projectEmployment").length) {
      var graphGradient = document.getElementById("projectEmployment").getContext('2d');
      var graphGradient2 = document.getElementById("projectEmployment").getContext('2d');
      var saleGradientBg = graphGradient.createLinearGradient(5, 0, 5, 100);
      saleGradientBg.addColorStop(0, 'rgba(42, 33, 186, 0.2)');
      saleGradientBg.addColorStop(1, 'rgba(42, 33, 186, 0)');
      var saleGradientBg2 = graphGradient2.createLinearGradient(100, 0, 50, 150);
      saleGradientBg2.addColorStop(0, 'rgba(0, 205, 255, 0.2)');
      saleGradientBg2.addColorStop(1, 'rgba(0, 205, 255, 0)');
      var salesTopData = {
          labels: ["SUN","sun", "MON", "mon", "TUE","tue", "WED", "wed", "THU", "thu", "FRI", "fri", "SAT"],
          datasets: [{
              label: 'Project',
              data: [50, 110, 60, 290, 200, 115, 130, 170, 90, 210, 240, 280, 200],
              backgroundColor: saleGradientBg,
              borderColor: [
                  '#2A21BA',
              ],
              borderWidth: 1.5,
              fill: true, // 3: no fill
              pointBorderWidth: 1,
              pointRadius: [4, 4, 4, 4, 4,4, 4, 4, 4, 4,4, 4, 4],
              pointHoverRadius: [4, 4, 4, 4, 4,4, 4, 4, 4, 4,4, 4, 4],
              pointBackgroundColor: ['#1F3BB3','#1F3BB3','#1F3BB3','#1F3BB3','#1F3BB3','#1F3BB3','#1F3BB3','#1F3BB3','#1F3BB3','#1F3BB3','#1F3BB3','#1F3BB3','#1F3BB3','#1F3BB3'],
              pointBorderColor: ['#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff'],
          },{
            label: 'Bench',
            data: [30, 150, 190, 250, 120, 150, 130, 20, 30, 15, 40, 95, 180],
            backgroundColor: saleGradientBg2,
            borderColor: [
                '#52CDFF',
            ],
            borderWidth: 1.5,
            fill: true, // 3: no fill
            pointBorderWidth: 1,
            pointRadius: [4, 4, 4, 4, 4,4, 4, 4, 4, 4,4, 4, 4],
            pointHoverRadius: [4, 4, 4, 4, 4,4, 4, 4, 4, 4,4, 4, 4],
            pointBackgroundColor: ['#00CDFF','#00CDFF','#00CDFF','#00CDFF','#00CDFF','#00CDFF','#00CDFF','#00CDFF','#00CDFF','#00CDFF','#00CDFF','#00CDFF','#00CDFF','#00CDFF'],
            pointBorderColor: ['#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff','#fff'],
        }]
      };
  
      var salesTopOptions = {
        responsive: true,
        maintainAspectRatio: true,
          scales: {
              yAxes: [{
                  gridLines: {
                      display: true,
                      drawBorder: false,
                      color:"#F0F0F0",
                      zeroLineColor: '#F0F0F0',
                  },
                  ticks: {
                    beginAtZero: false,
                    autoSkip: true,
                    maxTicksLimit: 4,
                    fontSize: 10,
                    color:"#6B778C"
                  }
              }],
              xAxes: [{
                gridLines: {
                    display: false,
                    drawBorder: false,
                },
                ticks: {
                  beginAtZero: false,
                  autoSkip: true,
                  maxTicksLimit: 7,
                  fontSize: 10,
                  color:"#6B778C"
                }
            }],
          },
          legend:false,
          legendCallback: function (chart) {
            var text = [];
            text.push('<div class="chartjs-legend"><ul>');
            for (var i = 0; i < chart.data.datasets.length; i++) {
              text.push('<li>');
              text.push('<span style="background-color:' + chart.data.datasets[i].borderColor + '">' + '</span>');
              text.push(chart.data.datasets[i].label);
              text.push('</li>');
            }
            text.push('</ul></div>');
            return text.join("");
          },
          elements: {
              line: {
                  tension: 0.4,
              }
          },
          tooltips: {
              backgroundColor: 'rgba(31, 59, 179, 1)',
          }
      }
      var salesTop = new Chart(graphGradient, {
          type: 'line',
          data: salesTopData,
          options: salesTopOptions
      });
      document.getElementById('projectEmploymentlegend').innerHTML = salesTop.generateLegend();
    }
    if ($("#doughnutCharthr").length) {
      var doughnutChartCanvas = $("#doughnutCharthr").get(0).getContext("2d");
      var doughnutPieData = {
        labels: [
          'Developers',
          'Marketing',
          'Finance',
          'Designing'
        ],
        datasets: [{
          data: [50, 20, 20, 10],
          backgroundColor: [
            "#1F3BB3",
            "#00CDFF",
            "#F95F53",
            "#00AAB6"
          ],
          borderColor: [
            "#fff",
            "#fff",
            "#fff",
            "#fff"
          ],
        }],
  
      };
      var doughnutPieOptions = {
        cutoutPercentage: 60,
        animationEasing: "easeOutBounce",
        animateRotate: true,
        animateScale: false,
        responsive: true,
        maintainAspectRatio: true,
        showScale: true,
        legend: false,
        
        layout: {
          padding: {
            left: 0,
            right: 0,
            top: 0,
            bottom: 0
          }
        },
        tooltips: {
          backgroundColor: '#fff',
          titleFontSize: 14,
          titleFontColor: '#0B0F32',
          bodyFontColor: '#737F8B',
          bodyFontSize: 11,
          displayColors: false
        }
      };
      var doughnutChart = new Chart(doughnutChartCanvas, {
        type: 'doughnut',
        data: doughnutPieData,
        options: doughnutPieOptions
      });
    }
    if ($("#projectstatus").length) {
      var marketingOverviewChart = document.getElementById("projectstatus").getContext('2d');
      var marketingOverviewData = {
          labels: ["JAN","FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEP", "OCT", "NOV", "DEC"],
          datasets: [{
            label: 'Finished',
            data: [110, 220, 200, 190, 220, 110, 210, 110, 205, 202, 201, 150],
            backgroundColor: "#00CDFF",
            borderColor: [
                '#00CDFF',
            ],
            borderWidth: 0,
            fill: true, // 3: no fill
              
          },{
            label: 'Pending',
            data: [215, 290, 210, 250, 290, 230, 290, 210, 280, 220, 190, 300],
            backgroundColor: "#1E3BB3",
            borderColor: [
                '#1E3BB3',
            ],
            borderWidth: 0,
            fill: true, // 3: no fill
        }]
      };
  
      var marketingOverviewOptions = {
        responsive: true,
        maintainAspectRatio: false,
        scales: {
          yAxes: [{
              gridLines: {
                  display: true,
                  drawBorder: false,
                  color:"#F0F0F0",
                  zeroLineColor: '#F0F0F0',
              },
              ticks: {
                beginAtZero: true,
                autoSkip: true,
                maxTicksLimit: 5,
                fontSize: 10,
                color:"#6B778C"
              }
          }],
          xAxes: [{
            stacked: true,
            barPercentage: 0.35,
            gridLines: {
                display: false,
                drawBorder: false,
            },
            ticks: {
              beginAtZero: false,
              autoSkip: true,
              maxTicksLimit: 12,
              fontSize: 10,
              color:"#6B778C"
            }
        }],
        },
        legend:false,
        legendCallback: function (chart) {
          var text = [];
          text.push('<div class="chartjs-legend"><ul>');
          for (var i = 0; i < chart.data.datasets.length; i++) {
            text.push('<li class="text-muted text-small">');
            text.push('<span style="background-color:' + chart.data.datasets[i].borderColor + '">' + '</span>');
            text.push(chart.data.datasets[i].label);
            text.push('</li>');
          }
          text.push('</ul></div>');
          return text.join("");
        },
        
        elements: {
            line: {
                tension: 0.4,
            }
        },
        tooltips: {
          backgroundColor: 'rgba(31, 59, 179, 1)',
        }
      }
      var marketingOverview = new Chart(marketingOverviewChart, {
          type: 'bar',
          data: marketingOverviewData,
          options: marketingOverviewOptions
      });
      document.getElementById('projectstatus-legend').innerHTML = marketingOverview.generateLegend();
    }
     
    if ($('#acceptedApplications').length) {

      var bar = new ProgressBar.Circle(acceptedApplications, {
        color: '#fff',
        svgStyle: {
          strokeLinecap: 'round',
        },
        // This has to be the same size as the maximum width to
        // prevent clipping
        strokeWidth: 10,
        trailWidth: 8,
        easing: 'easeInOut',
        duration: 1400,
        trailColor: 'rgba(255,255,255, .2)',
        text: { 
          autoStyleContainer: false
        },
        from: {
          color: '#00CDFF',
          width: 10,
          radius: 100,
        },
        to: {
          color: '#00CDFF',
          width: 10
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
  
      bar.text.style.fontSize = '1.875rem';
      bar.text.style.fontWeight = 'bold';
      bar.animate(.30); // Number from 0.0 to 1.0
    }
    if ($('#rejectedApplications').length) {
      var bar = new ProgressBar.Circle(rejectedApplications, {
        color: '#fff',
        svgStyle: {
          strokeLinecap: 'round',
        },
        // This has to be the same size as the maximum width to
        // prevent clipping
        strokeWidth: 10,
        trailWidth: 8,
        easing: 'easeInOut',
        duration: 1400,
        trailColor: 'rgba(255,255,255, .2)',
        text: { 
          autoStyleContainer: false
        },
        from: {
          color: '#FFFFFF',
          width: 10,
          radius: 100,
        },
        to: {
          color: '#FFFFFF',
          width: 10
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
  
      bar.text.style.fontSize = '1.875rem';
      bar.text.style.fontWeight = 'bold';
      bar.animate(.20); // Number from 0.0 to 1.0
    }

    if ($.cookie('staradmin2-pro-banner')!="true") {
      document.querySelector('#proBanner').classList.add('d-flex');
      document.querySelector('.navbar').classList.remove('fixed-top');
    }
    else {
      document.querySelector('#proBanner').classList.add('d-none');
      document.querySelector('.navbar').classList.add('fixed-top');
    }
    
    if ($( ".navbar" ).hasClass( "fixed-top" )) {
      document.querySelector('.page-body-wrapper').classList.remove('pt-0');
      document.querySelector('.navbar').classList.remove('pt-5');
    }
    else {
      document.querySelector('.page-body-wrapper').classList.add('pt-0');
      document.querySelector('.navbar').classList.add('pt-5');
      document.querySelector('.navbar').classList.add('mt-3');
      
    }
    document.querySelector('#bannerClose').addEventListener('click',function() {
      document.querySelector('#proBanner').classList.add('d-none');
      document.querySelector('#proBanner').classList.remove('d-flex');
      document.querySelector('.navbar').classList.remove('pt-5');
      document.querySelector('.navbar').classList.add('fixed-top');
      document.querySelector('.page-body-wrapper').classList.add('proBanner-padding-top');
      document.querySelector('.navbar').classList.remove('mt-3');
      var date = new Date();
      date.setTime(date.getTime() + 24 * 60 * 60 * 1000); 
      $.cookie('staradmin2-pro-banner', "true", { expires: date });
    });
    
  });
  // iconify.load('icons.svg').then(function() {
  //   iconify(document.querySelector('.my-cool.icon'));
  // });

  if ($("#realTimeUserAnalytic").length) {
    var realTimeUserAnalyticChart = document.getElementById("realTimeUserAnalytic").getContext('2d');
    var realTimegradient = realTimeUserAnalyticChart.createLinearGradient(1, 0, 1, 70);
    realTimegradient.addColorStop(1, 'rgba(30, 59, 179, 0.1)');
    realTimegradient.addColorStop(0, 'rgba(30, 59, 179, 0.8)');
    var realTimeUserAnalyticData = {
        labels: ["jan","feb", "mar", "apr", "may", "jun", "july", "aug", "sep", "oct"],
        datasets: [{
          label: 'Last week',
          data: [0, 10, 9, 16, 15, 17, 16, 18, 14, 25],
          backgroundColor: realTimegradient,
          borderWidth: 0,
          pointBorderWidth:0,
          borderColor: [
              '#1E3BB3',
          ],
          fill: true, // 3: no fill   
        }]
    };
    var realTimeUserAnalyticOptions = {
      responsive: true,
      maintainAspectRatio: false,
      scales: {
        yAxes: [{
          display: false,
            gridLines: {
                display: true,
                drawBorder: false,
                color:"#F0F0F0",
                zeroLineColor: '#F0F0F0',
            },
            ticks: {
              beginAtZero: true,
              autoSkip: true,
              maxTicksLimit: 6,
              fontSize: 10,
              color:"#6B778C"
            }
        }],
        xAxes: [{
          stacked: false,
          barPercentage: 0.5,
          display: false,
          gridLines: {
              display: false,
              drawBorder: false,
          },
          ticks: {
            beginAtZero: false,
            autoSkip: true,
            maxTicksLimit: 25,
            fontSize: 10,
            color:"#6B778C"
          }
      }],
      },
      legend:false,
      
      elements: {
          line: {
              tension: 0.4,
          },
          point:{
            radius: 0
          }
      },
      tooltips: {
        backgroundColor: 'rgba(31, 59, 179, 1)',
      }
    }
    var realTimeUserAnalyticChart = new Chart(realTimeUserAnalytic, {
        type: 'line',
        data: realTimeUserAnalyticData,
        options: realTimeUserAnalyticOptions
    });
  }
  if ($("#totalEarningsAnalytic").length) {
    var totalEarningsAnalyticChart = document.getElementById("totalEarningsAnalytic").getContext('2d');
    var totalEarningsgradient = totalEarningsAnalyticChart.createLinearGradient(1, 0, 1, 70);
    totalEarningsgradient.addColorStop(1, 'rgba(0, 170, 183, 0.1)');
    totalEarningsgradient.addColorStop(0, 'rgba(0, 170, 183, 0.8)');
    var totalEarningsAnalyticData = {
        labels: ["jan","feb", "mar", "apr", "may", "jun", "july", "aug", "sep", "oct"],
        datasets: [{
          label: 'Last week',
          data: [0, 10, 9, 16, 15, 17, 16, 18, 14, 25],
          backgroundColor: totalEarningsgradient,
          borderWidth: 0,
          pointBorderWidth:0,
          borderColor: [
              '#00AAB7',
          ],
          fill: true, // 3: no fill   
        }]
    };
    var totalEarningsAnalyticOptions = {
      responsive: true,
      maintainAspectRatio: false,
      scales: {
        yAxes: [{
          display: false,
            gridLines: {
                display: true,
                drawBorder: false,
                color:"#F0F0F0",
                zeroLineColor: '#F0F0F0',
            },
            ticks: {
              beginAtZero: true,
              autoSkip: true,
              maxTicksLimit: 6,
              fontSize: 10,
              color:"#6B778C"
            }
        }],
        xAxes: [{
          stacked: false,
          barPercentage: 0.5,
          display: false,
          gridLines: {
              display: false,
              drawBorder: false,
          },
          ticks: {
            beginAtZero: false,
            autoSkip: true,
            maxTicksLimit: 25,
            fontSize: 10,
            color:"#6B778C"
          }
      }],
      },
      legend:false,
      
      elements: {
          line: {
              tension: 0.4,
          },
          point:{
            radius: 0
          }
      },
      tooltips: {
        backgroundColor: 'rgba(31, 59, 179, 1)',
      }
    }
    var totalEarningsAnalyticChart = new Chart(totalEarningsAnalytic, {
        type: 'line',
        data: totalEarningsAnalyticData,
        options: totalEarningsAnalyticOptions
    });
  }
  if ($("#impressionAnalytic").length) {
    var impressionAnalyticChart = document.getElementById("impressionAnalytic").getContext('2d');
    var impressiongradient = impressionAnalyticChart.createLinearGradient(1, 0, 1, 70);
    impressiongradient.addColorStop(1, 'rgba(77, 167, 97, 0.1)');
    impressiongradient.addColorStop(0, 'rgba(77, 167, 97, 0.8)');
    var impressionAnalyticData = {
        labels: ["jan","feb", "mar", "apr", "may", "jun", "july", "aug", "sep", "oct"],
        datasets: [{
          label: 'Last week',
          data: [0, 10, 9, 16, 15, 17, 16, 18, 14, 25],
          backgroundColor: impressiongradient,
          borderWidth: 0,
          pointBorderWidth:0,
          borderColor: [
              '#4DA761',
          ],
          fill: true, // 3: no fill   
        }]
    };
    var impressionAnalyticOptions = {
      responsive: true,
      maintainAspectRatio: false,
      scales: {
        yAxes: [{
          display: false,
            gridLines: {
                display: true,
                drawBorder: false,
                color:"#F0F0F0",
                zeroLineColor: '#F0F0F0',
            },
            ticks: {
              beginAtZero: true,
              autoSkip: true,
              maxTicksLimit: 6,
              fontSize: 10,
              color:"#6B778C"
            }
        }],
        xAxes: [{
          stacked: false,
          barPercentage: 0.5,
          display: false,
          gridLines: {
              display: false,
              drawBorder: false,
          },
          ticks: {
            beginAtZero: false,
            autoSkip: true,
            maxTicksLimit: 25,
            fontSize: 10,
            color:"#6B778C"
          }
      }],
      },
      legend:false,
      
      elements: {
          line: {
              tension: 0.4,
          },
          point:{
            radius: 0
          }
      },
      tooltips: {
        backgroundColor: 'rgba(31, 59, 179, 1)',
      }
    }
    var impressionAnalyticChart = new Chart(impressionAnalytic, {
        type: 'line',
        data: impressionAnalyticData,
        options: impressionAnalyticOptions
    });
  }
  if ($("#siteIncomeAnalytic").length) {
    var siteIncomeAnalyticChart = document.getElementById("siteIncomeAnalytic").getContext('2d');
    var siteIncomegradient = siteIncomeAnalyticChart.createLinearGradient(1, 0, 1, 70);
    siteIncomegradient.addColorStop(1, 'rgba(249, 95, 83, 0.1)');
    siteIncomegradient.addColorStop(0, 'rgba(249, 95, 83, 0.8)');
    var siteIncomeAnalyticData = {
        labels: ["jan","feb", "mar", "apr", "may", "jun", "july", "aug", "sep", "oct"],
        datasets: [{
          label: 'Last week',
          data: [0, 10, 9, 16, 15, 17, 16, 18, 14, 25],
          backgroundColor: siteIncomegradient,
          borderWidth: 0,
          pointBorderWidth:0,
          borderColor: [
              '#F95F53',
          ],
          fill: true, // 3: no fill   
        }]
    };
    var siteIncomeAnalyticOptions = {
      responsive: true,
      maintainAspectRatio: false,
      scales: {
        yAxes: [{
          display: false,
            gridLines: {
                display: true,
                drawBorder: false,
                color:"#F0F0F0",
                zeroLineColor: '#F0F0F0',
            },
            ticks: {
              beginAtZero: true,
              autoSkip: true,
              maxTicksLimit: 6,
              fontSize: 10,
              color:"#6B778C"
            }
        }],
        xAxes: [{
          stacked: false,
          barPercentage: 0.5,
          display: false,
          gridLines: {
              display: false,
              drawBorder: false,
          },
          ticks: {
            beginAtZero: false,
            autoSkip: true,
            maxTicksLimit: 25,
            fontSize: 10,
            color:"#6B778C"
          }
      }],
      },
      legend:false,
      
      elements: {
          line: {
              tension: 0.4,
          },
          point:{
            radius: 0
          }
      },
      tooltips: {
        backgroundColor: 'rgba(31, 59, 179, 1)',
      }
    }
    var siteIncomeAnalyticChart = new Chart(siteIncomeAnalytic, {
        type: 'line',
        data: siteIncomeAnalyticData,
        options: siteIncomeAnalyticOptions
    });
  }
  if ($("#doughnutChartAnalytic").length) {
    var doughnutChartCanvasAnalytic = $("#doughnutChartAnalytic").get(0).getContext("2d");
    var doughnutPieDataAnalytic = {
      datasets: [{
        data: [50, 20, 30],
        backgroundColor: [
          "#1F3BB3",
          "#00CDFF",
          "#F95F53",
        ],
        borderColor: [
          "#fff",
          "#fff",
          "#fff",
        ],
      }],

      // These labels appear in the legend and in the tooltips when hovering different arcs
      labels: [
        'Admin dashboard',
        'Website design',
        'Mobile app design',
      ]
    };
    var doughnutPieOptionsAnalytic = {
      cutoutPercentage: 50,
      animationEasing: "easeOutBounce",
      animateRotate: true,
      animateScale: false,
      responsive: true,
      maintainAspectRatio: true,
      showScale: true,
      legend: false,
      legendCallback: function (chart) {
        var text = [];
        text.push('<div class="chartjs-legend"><ul>');
        for (var i = 0; i < chart.data.datasets[0].data.length; i++) {
          
          if (chart.data.labels[i]) {
            text.push(chart.data.datasets[0].data[i] + '%');
          }


          text.push('<li><span style="background-color:' + chart.data.datasets[0].backgroundColor[i] + '">');
          text.push('</span>');
          
          if (chart.data.labels[i]) {
            text.push(chart.data.labels[i]);
          }
          text.push('</li>');
        }
        text.push('</div></ul>');
        return text.join("");
      },
      
      layout: {
        padding: {
          left: 0,
          right: 0,
          top: 0,
          bottom: 0
        }
      },
      tooltips: {
        callbacks: {
          title: function(tooltipItem, data) {
            return data['labels'][tooltipItem[0]['index']];
          },
          label: function(tooltipItem, data) {
            return data['datasets'][0]['data'][tooltipItem['index']];
          }
        },
          
        backgroundColor: '#fff',
        titleFontSize: 14,
        titleFontColor: '#0B0F32',
        bodyFontColor: '#737F8B',
        bodyFontSize: 11,
        displayColors: false
      }
    };
    var doughnutChartAnalytic = new Chart(doughnutChartCanvasAnalytic, {
      type: 'doughnut',
      data: doughnutPieDataAnalytic,
      options: doughnutPieOptionsAnalytic
    });
    document.getElementById('doughnut-chart-legend-Analytic').innerHTML = doughnutChartAnalytic.generateLegend();
  }

  if ($("#realtimestatisticsAnalytic").length) {
    var realtimestatisticsAnalyticChart = document.getElementById("realtimestatisticsAnalytic").getContext('2d');
    var realtimestatisticsAnalyticData = {
        labels: ["Jan","Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"],
        datasets: [{
          label: 'Last week',
          data: [125, 169, 60, 140, 100, 170, 50, 80, 240, 140, 80, 160],
          backgroundColor: "#1E3BB3",
          borderColor: [
              '#1E3BB3',
          ],
          borderWidth: 0,
          fill: true, // 3: no fill
            
        },
        {
          label: 'Last week',
          
          data: [200, 290, 220, 180, 200, 250, 120, 170, 290, 210, 170, 210],
          backgroundColor: "#E3E9FF",
          borderColor: [
              '#E3E9FF',
          ],
          borderWidth: 0,
          fill: true, // 3: no fill
            
        }]
    };

    var realtimestatisticsAnalyticOptions = {
      responsive: true,
      maintainAspectRatio: false,
      scales: {
        yAxes: [{
            gridLines: {
                display: true,
                drawBorder: false,
                color:"#F0F0F0",
                zeroLineColor: '#F0F0F0',
            },
            ticks: {
              beginAtZero: true,
              autoSkip: true,
              maxTicksLimit: 6,
              fontSize: 10,
              color:"#6B778C"
            }
        }],
        xAxes: [{
          stacked: true,
          barPercentage: 0.4,
          gridLines: {
              display: false,
              drawBorder: false,
          },
          ticks: {
            beginAtZero: false,
            autoSkip: true,
            maxTicksLimit: 12,
            fontSize: 10,
            color:"#6B778C"
          }
      }],
      },
      legend:false,
      
      elements: {
          line: {
              tension: 0.4,
          }
      },
      tooltips: {
        backgroundColor: 'rgba(31, 59, 179, 1)',
      }
    }
    var realtimestatisticsAnalytic = new Chart(realtimestatisticsAnalyticChart, {
        type: 'bar',
        data: realtimestatisticsAnalyticData,
        options: realtimestatisticsAnalyticOptions
    });
  }

  if ($('#totalVisitorsanalytic').length) {
    var bar = new ProgressBar.Circle(totalVisitorsanalytic, {
      color: '#fff',
      // This has to be the same size as the maximum width to
      // prevent clipping
      strokeWidth: 15,
      trailWidth: 15, 
      easing: 'easeInOut',
      duration: 1400,
      text: {
        autoStyleContainer: false
      },
      from: {
        color: '#00CDFF',
        width: 15
      },
      to: {
        color: '#00CDFF',
        width: 15
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

    bar.text.style.fontSize = '0rem';
    bar.animate(.64); // Number from 0.0 to 1.0
  }
  if ($('#visitperdayAnalytic').length) {
    var bar = new ProgressBar.Circle(visitperdayAnalytic, {
      color: '#fff',
      // This has to be the same size as the maximum width to
      // prevent clipping
      strokeWidth: 15,
      trailWidth: 15,
      easing: 'easeInOut',
      duration: 1400,
      text: {
        autoStyleContainer: false
      },
      from: {
        color: '#1E3BB3',
        width: 15
      },
      to: {
        color: '#1E3BB3',
        width: 15
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

    bar.text.style.fontSize = '0rem';
    bar.animate(.34); // Number from 0.0 to 1.0
  }
  if ($("#performanceAnalytic").length) {
    var performanceAnalyticChart = document.getElementById("performanceAnalytic").getContext('2d');
    var performanceAnalyticgradient = performanceAnalyticChart.createLinearGradient(10, 10, 1, 160);
    performanceAnalyticgradient.addColorStop(1, 'rgba(0, 170, 183, 0)');
    performanceAnalyticgradient.addColorStop(0, 'rgba(0, 170, 183, 0.6)');
    var performanceAnalyticData = {
        labels: ["one","two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen"],
        datasets: [{
          label: 'Last week',
          data: [30, 20, 25, 22, 35, 18, 22, 20, 34, 17, 24, 22, 36],
          borderWidth: 0,
          pointBorderWidth:0,
          borderColor: [
              '#00AAB7',
          ],
          fill: false, // 3: no fill   
        },
        {
          label: 'Last week',
          data: [28, 18, 23, 20, 33, 16, 20, 18, 32, 15, 22, 20, 34],
          backgroundColor: performanceAnalyticgradient,
          borderWidth: 0,
          pointBorderWidth:0,
          fill: true, // 3: no fill   
        }]
    };
    var performanceAnalyticOptions = {
      responsive: true,
      maintainAspectRatio: false,
      scales: {
        yAxes: [{
          display: false,
            gridLines: {
                display: true,
                drawBorder: false,
                color:"#F0F0F0",
                zeroLineColor: '#F0F0F0',
            },
            ticks: {
              beginAtZero: true,
              autoSkip: true,
              maxTicksLimit: 38,
              fontSize: 10,
              color:"#6B778C"
            }
        }],
        xAxes: [{
          stacked: false,
          barPercentage: 0.5,
          display: false,
          gridLines: {
              display: false,
              drawBorder: false,
          },
          ticks: {
            beginAtZero: false,
            autoSkip: true,
            maxTicksLimit: 38,
            fontSize: 10,
            color:"#6B778C"
          }
      }],
      },
      legend:false,
      
      elements: {
          line: {
              tension: 0,
          },
          point:{
            radius: 0
          }
      },
      tooltips: {
        backgroundColor: 'rgba(31, 59, 179, 1)',
      }
    }
    var performanceAnalyticChart = new Chart(performanceAnalytic, {
        type: 'line',
        data: performanceAnalyticData,
        options: performanceAnalyticOptions
    });
  }


  if ($("#modernRevenueGrowth").length) {
    var modernRevenueGrowthChart = document.getElementById("modernRevenueGrowth").getContext('2d');
    var modernRevenueGrowthData = {
        labels: ["Jan","Feb", "Mar", "Apr", "May", "Jun", "Jul"],
        datasets: [{
          label: 'Last week',
          data: [50, 75, 100, 60, 70, 45, 90],
          backgroundColor: "#00CDFF",
          borderColor: [
              '#00CDFF',
          ],
          borderWidth: 0,
          fill: true, // 3: no fill
            
        }]
    };

    var modernRevenueGrowthOptions = {
      responsive: true,
      maintainAspectRatio: false,
      scales: {
        yAxes: [{
            gridLines: {
                display: true,
                drawBorder: false,
                color:"#F0F0F0",
                zeroLineColor: '#F0F0F0',
            },
            ticks: {
              beginAtZero: true,
              autoSkip: true,
              maxTicksLimit: 6,
              fontSize: 10,
              color:"#6B778C"
            }
        }],
        xAxes: [{
          stacked: false,
          barPercentage: 0.4,
          gridLines: {
              display: false,
              drawBorder: false,
          },
          ticks: {
            beginAtZero: false,
            autoSkip: true,
            maxTicksLimit: 12,
            fontSize: 10,
            color:"#6B778C"
          }
      }],
      },
      legend:false,
      
      elements: {
          line: {
              tension: 0.4,
          }
      },
      tooltips: {
        backgroundColor: 'rgba(31, 59, 179, 1)',
      }
    }
    var modernRevenueGrowth = new Chart(modernRevenueGrowthChart, {
        type: 'bar',
        data: modernRevenueGrowthData,
        options: modernRevenueGrowthOptions
    });
  }

  if ($("#modernBubble").length) {
    var modernBubbleChart = document.getElementById("modernBubble").getContext('2d');
    var modernBubbleData = {
        labels: ["Jan","Feb", "Mar", "Apr", "May", "Jun", "Jul"],
        datasets: [{
          label: 'Money send',
          data: [{
            x: 10,
            y: 100,
            r: 10
          }, {
            x: 20,
            y: 500,
            r: 15
          }, {
            x: 40,
            y: 100,
            r: 10
          }, {
            x: 55,
            y: 200,
            r: 10
          }, {
            x: 70,
            y: 500,
            r: 10
          }, {
            x: 0,
            y: 600,
            r: 0
          }],
          backgroundColor: 'rgb(30,59,179)'
        },{
          label: 'Money Received',
          data: [{
            x: 10,
            y: 300,
            r: 5
          }, {
            x: 30,
            y: 400,
            r: 5
          }, {
            x: 60,
            y: 410,
            r: 10
          }, {
            x: 100,
            y: 370,
            r: 5
          }, {
            x: 110,
            y: 0,
            r: 0
          }],
          backgroundColor: 'rgb(99,171,253)',
        }]
    };

    var modernBubbleOptions = {
      responsive: true,
      maintainAspectRatio: false,
      scales: {
        yAxes: [{
            gridLines: {
                display: true,
                drawBorder: false,
                color:"#F0F0F0",
                zeroLineColor: '#F0F0F0',
            },
            ticks: {
              beginAtZero: true,
              autoSkip: true,
              maxTicksLimit: 10,
              fontSize: 10,
              color:"#6B778C"
            }
        }],
        xAxes: [{
          stacked: false,
          barPercentage: 0.4,
          gridLines: {
              display: false,
              drawBorder: false,
          },
          ticks: {
            beginAtZero: false,
            autoSkip: true,
            maxTicksLimit: 10,
            fontSize: 10,
            color:"#6B778C"
          }
      }],
      },
      legend:false,
      legendCallback: function (chart) {
        var text = [];
        text.push('<div class="chartjs-legend"><ul>');
        for (var i = 0; i < chart.data.datasets.length; i++) {
          console.log(chart.data.datasets[i]); // see what's inside the obj.
          text.push('<li class="text-dark text-small">');
          text.push('<span style="background-color:' + chart.data.datasets[i].backgroundColor + '">' + '</span>');
          text.push(chart.data.datasets[i].label);
          text.push('</li>');
        }
        text.push('</ul></div>');
        return text.join("");
      },
      
      elements: {
          line: {
              tension: 0.4,
          }
      },
      tooltips: {
        backgroundColor: 'rgba(31, 59, 179, 1)',
      }
    }
    var modernBubble = new Chart(modernBubbleChart, {
        type: 'bubble',
        data: modernBubbleData,
        options: modernBubbleOptions
    });
    document.getElementById('bubble-chart-legend').innerHTML = modernBubble.generateLegend();
  }
  if ($("#moneyFlow").length) {
    var moneyFlowChart = document.getElementById("moneyFlow").getContext('2d');
    var moneyFlowgradient = moneyFlowChart.createLinearGradient(10, 10, 1, 160);
    moneyFlowgradient.addColorStop(1, 'rgba(30, 59, 179, 0)');
    moneyFlowgradient.addColorStop(0, 'rgba(30, 59, 179, 0.3)');
    var moneyFlowData = {
        labels: ["jan","feb", "mar", "apr", "may", "jun", "july", "aug", "sep", "oct", "nov", "dec"],
        datasets: [{
          label: 'Last week',
          data: [20000, 50000, 30000, 80000, 60000, 55000, 45000, 60000, 35000, 50000, 55000, 40000],
          backgroundColor: moneyFlowgradient,
          borderWidth: 2,
          pointBorderWidth:0,
          borderColor: [
              '#1E3BB3',
          ],
          fill: true, // 3: no fill   
        }]
    };
    var moneyFlowOptions = {
      responsive: true,
      maintainAspectRatio: false,
      scales: {
        yAxes: [{
          display: true,
            gridLines: {
                display: true,
                drawBorder: false,
                color:"#F0F0F0",
                zeroLineColor: '#F0F0F0',
            },
            ticks: {
              beginAtZero: true,
              autoSkip: true,
              maxTicksLimit: 6,
              fontSize: 10,
              color:"#6B778C"
            }
        }],
        xAxes: [{
          stacked: false,
          barPercentage: 0.5,
          display: true,
          gridLines: {
              display: false,
              drawBorder: false,
          },
          ticks: {
            beginAtZero: true,
            autoSkip: true,
            maxTicksLimit: 12,
            fontSize: 10,
            color:"#6B778C"
          }
      }],
      },
      legend:false,
      
      elements: {
          line: {
              tension: 0.5,
          },
          point:{
            radius: 0
          }
      },
      tooltips: {
        backgroundColor: 'rgba(31, 59, 179, 1)',
      }
    }
    var moneyFlowChart = new Chart(moneyFlow, {
        type: 'line',
        data: moneyFlowData,
        options:moneyFlowOptions
    });
  }
  if ($("#modernChartliability").length) {
    var modernChartliabilityCanvas = $("#modernChartliability").get(0).getContext("2d");
    var modernChartliabilityData = {
      datasets: [{
        data: [50, 20, 30],
        backgroundColor: [
          "#4DA761",
          "#00CDFF",
          "#EE5E51",
        ],
        borderColor: [
          "#fff",
          "#fff",
          "#fff",
        ],
      }],

      // These labels appear in the legend and in the tooltips when hovering different arcs
      labels: [
        'Current',
        'New',
        'Pending',
      ]
    };
    var modernChartliabilityOptions = {
      cutoutPercentage: 60,
      animationEasing: "easeOutBounce",
      animateRotate: true,
      animateScale: false,
      responsive: true,
      maintainAspectRatio: true,
      showScale: true,
      legend: false,
      legendCallback: function (chart) {
        var text = [];
        text.push('<div class="chartjs-legend"><ul>');
        for (var i = 0; i < chart.data.datasets[0].data.length; i++) {
          text.push('<li><span style="background-color:' + chart.data.datasets[0].backgroundColor[i] + '">');
          text.push('</span>');
          
          if (chart.data.labels[i]) {
            text.push(chart.data.labels[i]);
          }
          text.push('</li>');
        }
        text.push('</div></ul>');
        return text.join("");
      },
      
      layout: {
        padding: {
          left: 0,
          right: 0,
          top: 0,
          bottom: 0
        }
      },
      tooltips: {
        callbacks: {
          title: function(tooltipItem, data) {
            return data['labels'][tooltipItem[0]['index']];
          },
          label: function(tooltipItem, data) {
            return data['datasets'][0]['data'][tooltipItem['index']];
          }
        },
          
        backgroundColor: '#fff',
        titleFontSize: 14,
        titleFontColor: '#0B0F32',
        bodyFontColor: '#737F8B',
        bodyFontSize: 11,
        displayColors: false
      }
    };
    var modernChartliability = new Chart(modernChartliabilityCanvas, {
      type: 'doughnut',
      data: modernChartliabilityData,
      options: modernChartliabilityOptions
    });
    document.getElementById('doughnut-chart-legend-modern').innerHTML = modernChartliability.generateLegend();
  }
  if ($("#summarySales").length) {
    var summarySalesChart = document.getElementById("summarySales").getContext('2d');
    var summarySalesgradient = summarySalesChart.createLinearGradient(10, 10, 1, 160);
    var pointBg = ["rgba(255,255,255,0)","rgba(255,255,255,0)","rgba(255,255,255,0)","rgba(255,255,255,0)","rgba(255,255,255,0)","rgba(255,255,255,0)","#1E3BB3","rgba(255,255,255,0)","rgba(255,255,255,0)","rgba(255,255,255,0)","rgba(255,255,255,0)","rgba(255,255,255,0)"];
    var pointBorder = ["rgba(255,255,255,0)","rgba(255,255,255,0)","rgba(255,255,255,0)","rgba(255,255,255,0)","rgba(255,255,255,0)","rgba(255,255,255,0)","#fff","rgba(255,255,255,0)","rgba(255,255,255,0)","rgba(255,255,255,0)","rgba(255,255,255,0)","rgba(255,255,255,0)"];

    summarySalesgradient.addColorStop(1, 'rgba(30, 59, 179, 0)');
    summarySalesgradient.addColorStop(0, 'rgba(30, 59, 179, 0.3)');
    var summarySalesData = {
        labels: ["jan","feb", "mar", "apr", "may", "jun", "july", "aug", "sep", "oct", "nov", "dec"],
        datasets: [{
          label: 'Last week',
          data: [20000, 50000, 30000, 80000, 60000, 55000, 45000, 60000, 35000, 50000, 55000, 40000],
          backgroundColor: summarySalesgradient,
          borderWidth: 2,
          borderColor: [
              '#1E3BB3',
          ],
          fill: true, // 3: no fill  
          
          pointBackgroundColor:  pointBg,
          pointBorderColor: pointBorder,
          radius: 5,
          pointRadius: 5
        }]
    };
    var summarySalesOptions = {
      responsive: true,
      maintainAspectRatio: false,
      scales: {
        yAxes: [{
          display: true,
            gridLines: {
                display: true,
                drawBorder: false,
                color:"#F0F0F0",
                zeroLineColor: '#F0F0F0',
            },
            ticks: {
              beginAtZero: true,
              autoSkip: true,
              maxTicksLimit: 6,
              fontSize: 10,
              color:"#6B778C"
            }
        }],
        xAxes: [{
          stacked: false,
          barPercentage: 0.5,
          display: true,
          gridLines: {
              display: false,
              drawBorder: false,
          },
          ticks: {
            beginAtZero: true,
            autoSkip: true,
            maxTicksLimit: 12,
            fontSize: 10,
            color:"#6B778C"
          }
      }],
      },
      legend:false,
      
      elements: {
          line: {
              tension: 0.5,
          },
          point:{
            radius: 0
          }
      },
      tooltips: {
        backgroundColor: 'rgba(31, 59, 179, 1)',
      }
    }
    var summarySalesChart = new Chart(summarySales, {
        type: 'line',
        data: summarySalesData,
        options:summarySalesOptions
    });
  }
  if ($("#customerOverviewEcommerce").length) {
    var customerOverviewEcommerceCanvas = $("#customerOverviewEcommerce").get(0).getContext("2d");
    var customerOverviewEcommerceData = {
      datasets: [{
        data: [50, 20, 30],
        backgroundColor: [
          "#1E3BB3",
          "#00CDFF",
          "#00AAB7",
        ],
        borderColor: [
          "#fff",
          "#fff",
          "#fff",
        ],
      }],

      // These labels appear in the legend and in the tooltips when hovering different arcs
      labels: [
        'Current',
        'New',
        'Retargeted',
      ]
    };
    var customerOverviewEcommerceOptions = {
      cutoutPercentage: 60,
      animationEasing: "easeOutBounce",
      animateRotate: true,
      animateScale: false,
      responsive: true,
      maintainAspectRatio: true,
      showScale: true,
      legend: false,
      legendCallback: function (chart) {
        var text = [];
        text.push('<div class="chartjs-legend"><ul>');
        for (var i = 0; i < chart.data.datasets[0].data.length; i++) {
          text.push('<li><span style="background-color:' + chart.data.datasets[0].backgroundColor[i] + '">');
          text.push('</span>');
          
          if (chart.data.labels[i]) {
            text.push(chart.data.labels[i]);
          }
          text.push('</li>');
        }
        text.push('</div></ul>');
        return text.join("");
      },
      
      layout: {
        padding: {
          left: 0,
          right: 0,
          top: 0,
          bottom: 0
        }
      },
      tooltips: {
        callbacks: {
          title: function(tooltipItem, data) {
            return data['labels'][tooltipItem[0]['index']];
          },
          label: function(tooltipItem, data) {
            return data['datasets'][0]['data'][tooltipItem['index']];
          }
        },
          
        backgroundColor: '#fff',
        titleFontSize: 14,
        titleFontColor: '#0B0F32',
        bodyFontColor: '#737F8B',
        bodyFontSize: 11,
        displayColors: false
      }
    };
    var customerOverviewEcommerce = new Chart(customerOverviewEcommerceCanvas, {
      type: 'doughnut',
      data: customerOverviewEcommerceData,
      options: customerOverviewEcommerceOptions
    });
    document.getElementById('customerOverviewEcommerce-legend').innerHTML = customerOverviewEcommerce.generateLegend();
  }
  if ($("#totalSalesByUnit").length) {
    var totalSalesByUnitCanvas = $("#totalSalesByUnit").get(0).getContext("2d");
    var totalSalesByUnitData = {
      datasets: [{
        data: [20, 55, 25],
        backgroundColor: [
          "#4DA761",
          "#F95F53",
          "#00CDFF",
        ],
        borderColor: [
          "#4DA761",
          "#F95F53",
          "#00CDFF",
        ],
      }],

      // These labels appear in the legend and in the tooltips when hovering different arcs
      labels: [
        'Online',
        'Offline',
        'Marketing',
      ]
    };
    var totalSalesByUnitOptions = {
      cutoutPercentage: 0,
      animationEasing: "easeOutBounce",
      animateRotate: true,
      animateScale: false,
      responsive: true,
      maintainAspectRatio: true,
      showScale: true,
      legend: false,
      legendCallback: function (chart) {
        var text = [];
        text.push('<div class="chartjs-legend"><ul>');
        for (var i = 0; i < chart.data.datasets[0].data.length; i++) {
          text.push('<li><span style="background-color:' + chart.data.datasets[0].backgroundColor[i] + '">');
          text.push('</span>');
          
          if (chart.data.labels[i]) {
            text.push(chart.data.labels[i]);
          }
          text.push('</li>');
        }
        text.push('</div></ul>');
        return text.join("");
      },
      
      layout: {
        padding: {
          left: 0,
          right: 0,
          top: 0,
          bottom: 0
        }
      },
      tooltips: {
        callbacks: {
          title: function(tooltipItem, data) {
            return data['labels'][tooltipItem[0]['index']];
          },
          label: function(tooltipItem, data) {
            return data['datasets'][0]['data'][tooltipItem['index']];
          }
        },
          
        backgroundColor: '#fff',
        titleFontSize: 14,
        titleFontColor: '#0B0F32',
        bodyFontColor: '#737F8B',
        bodyFontSize: 11,
        displayColors: false
      }
    };
    var totalSalesByUnit = new Chart(totalSalesByUnitCanvas, {
      type: 'pie',
      data: totalSalesByUnitData,
      options: totalSalesByUnitOptions
    });
    document.getElementById('totalSalesByUnit-legend').innerHTML = totalSalesByUnit.generateLegend();
  }
  if ($('#carouselExampleControls').length) {
    var myCarousel = document.querySelector('#carouselExampleControls')
    var carousel = new bootstrap.Carousel(myCarousel)
  }
  if ($("#incomeExpences").length) {
    var incomeExpencesChart = document.getElementById("incomeExpences").getContext('2d');
    var incomeExpencesData = {
        labels: ["Jan","Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"],
        datasets: [{
          label: 'Income',
          data: [125, 169, 60, 140, 100, 170, 50, 80, 240, 140, 80, 160],
          backgroundColor: "#1E3BB3",
          borderColor: [
              '#1E3BB3',
          ],
          borderWidth: 0,
          fill: true, // 3: no fill
            
        },
        {
          label: 'Expense',
          
          data: [200, 290, 220, 180, 200, 250, 120, 170, 290, 210, 170, 210],
          backgroundColor: "#00CDFF",
          borderColor: [
              '#00CDFF',
          ],
          borderWidth: 0,
          fill: true, // 3: no fill
            
        }]
    };

    var incomeExpencesOptions = {
      responsive: true,
      maintainAspectRatio: false,
      scales: {
        yAxes: [{
            gridLines: {
                display: true,
                drawBorder: false,
                color:"#F0F0F0",
                zeroLineColor: '#F0F0F0',
            },
            ticks: {
              beginAtZero: true,
              autoSkip: true,
              maxTicksLimit: 6,
              fontSize: 10,
              color:"#6B778C"
            }
        }],
        xAxes: [{
          stacked: true,
          barPercentage: 0.4,
          gridLines: {
              display: false,
              drawBorder: false,
          },
          ticks: {
            beginAtZero: false,
            autoSkip: true,
            maxTicksLimit: 12,
            fontSize: 10,
            color:"#6B778C"
          }
      }],
      },
      legend:false,
      legendCallback: function (chart) {
        var text = [];
        text.push('<div class="chartjs-legend"><ul>');
        for (var i = 0; i < chart.data.datasets.length; i++) {
          text.push('<li>');
          text.push('<span style="background-color:' + chart.data.datasets[i].borderColor + '">' + '</span>');
          text.push(chart.data.datasets[i].label);
          text.push('</li>');
        }
        text.push('</ul></div>');
        return text.join("");
      },
      elements: {
          line: {
              tension: 0.4,
          }
      },
      tooltips: {
        backgroundColor: 'rgba(31, 59, 179, 1)',
      }
    }
    var incomeExpences = new Chart(incomeExpencesChart, {
        type: 'bar',
        data: incomeExpencesData,
        options: incomeExpencesOptions
    });
    document.getElementById('incomeExpences-legend').innerHTML = incomeExpences.generateLegend();
  }
})(jQuery);