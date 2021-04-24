sql.ErrNoRows是sentinel errors，当dao层query时遇到sql.ErrNoRows时，若不wrap就上抛，上层调用者无法知道错误的详细信息。
Go1.13后可以使用fmt.Errorf("... %w", err) 做封装