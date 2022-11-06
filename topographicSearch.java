class Solution {
    HashSet<Integer> hashset = new HashSet<>();
    ArrayList<ArrayList<Integer>> graph = new ArrayList<>();
    public int[] findOrder(int numCourses, int[][] prerequisites) {
        int[] result = new int[numCourses];
        if(prerequisites.length==0){
            for(int i =0; i<numCourses; i++){
                result[i]=i;
            }
            return result;
        }
        for(int i =0; i < numCourses; i++){
            graph.add(new ArrayList<Integer>());
        }
        for(int[] array : prerequisites){
            graph.get(array[0]).add(array[1]);
        }
        ArrayList<Integer> tempResult = new ArrayList<>();
        for(Integer i=0; i<numCourses; i++){
            result = traverse(graph,i, tempResult).stream().mapToInt(j -> j).toArray();
        }
        return result;
    }
    
    private ArrayList<Integer> 
        traverse(ArrayList<ArrayList<Integer>> graph, Integer element
                 , ArrayList<Integer> result){
        if(hashset.contains(element)){return result;}
        if(graph.get(element).isEmpty()){
            result.add(element);
            hashset.add(element);
            return result;
        }
        for(Integer i : graph.get(element)){
            return traverse(graph, i, result);
        }
        return result;
    }

}
