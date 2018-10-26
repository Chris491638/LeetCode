# @param {Integer[]} nums
# @param {Integer} target
# @return {Integer[]}
def two_sum(nums, target)
    # 236ms 35.71% chris
    # rs = []
    # (0..nums.size-1).each do |i|
    #     rs[0] = i
    #     if nums[i+1..-1].index(target-nums[i])   
    #         rs[1] = nums[i+1..-1].index(target-nums[i])+i+1
    #         break
    #     end
    # end  
    # rs
    
    #48ms 89.55% chris
    hash = {}
    nums.each_with_index {|e,i| hash[e]=i}
    nums.each_with_index do |e,i|
        if hash[target-e] && hash[target-e]!=i
           return [i,hash[target-e]] 
        end
    end
    
    # 44ms 100% other
    # hash = {}
    # nums.each_with_index do |num, idx|
    #     check = target - num
    #     if hash[check] && hash[check] != idx
    #         return [idx, hash[check]]
    #     else
    #         hash[num] = idx
    #     end
    # end
end