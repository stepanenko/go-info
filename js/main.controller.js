
function MainController() {
  this.forecast = '';
}


app.controller('MainController', ['$scope', 'forecast', function($scope, forecast) {
  forecast.success(function(data) {
    $scope.fiveDay = data;
  });
}]);

angular
  .module('app')
  .controller('MainController', MainController);
